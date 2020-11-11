package pipe

import (
	"fmt"
	"io"
	"os"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/gocarina/gocsv"
)

const (
	Pos = "positive"
	Neg = "negative"
)

type RawLabResult struct {
	Id     string `csv:"student_id"`
	Result string `csv:"result"`
}

func (r RawLabResult) IsValid() bool {
	return r.Id != "" &&
		r.Result != "" &&
		(r.Result == Pos || r.Result == Neg)
}

func (r RawLabResult) ToLabResult(patient types.Patient) (*lab.Result, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("Attempt to convert invalid result %v", r)
	} else if patient.Id() != r.Id {
		return nil, fmt.Errorf("Attempt to create result with mismatching keys, %s %s", patient.Id(), r.Id)
	}

	result := lab.MakeResult(patient, r.Result == Pos)
	return &result, nil
}

type RawPatient struct {
	Id     string `csv:"student_id"`
	Name   string `csv:"full_name"`
	Number string `csv:"cell_phone_number"`
}

func (r *RawPatient) IsValid() bool {
	return r.Id != "" &&
		r.Name != "" &&
		r.Number != ""
}

func LoadPatients(patientsIn io.Reader) ([]types.Patient, error) {
	rawPatients := []*RawPatient{}
	err := gocsv.Unmarshal(patientsIn, &rawPatients)
	if err != nil {
		return nil, err
	}

	patients := make([]types.Patient, len(rawPatients))
	for i, rp := range rawPatients {

		if !rp.IsValid() {
			return nil, fmt.Errorf("Error loading patients, example %v", rp)
		}

		phone := types.MakePhoneNumber(rp.Number)
		patients[i] = types.MakePatient(rp.Id, rp.Name, phone)
	}

	return patients, nil
}

func LoadLabResults(resultsIn io.Reader) ([]*RawLabResult, error) {

	rawResults := []*RawLabResult{}
	err := gocsv.Unmarshal(resultsIn, &rawResults)
	if err != nil {
		return nil, err
	}

	for _, rl := range rawResults {

		if !rl.IsValid() {
			return nil, fmt.Errorf("Error loading lab results, example %v", rl)
		}
	}

	return rawResults, nil
}

func HashJoin(rawResults []*RawLabResult, patients []types.Patient) ([]*lab.Result, error) {
	m := make(map[string]*RawLabResult)
	for _, rl := range rawResults {
		if _, ok := m[rl.Id]; ok {
			return nil, fmt.Errorf("Duplicate key in results %s", rl.Id)
		}
		m[rl.Id] = rl
	}

	results := make([]*lab.Result, len(patients))
	for i, patient := range patients {
		result := m[patient.Id()]
		if result == nil {
			return nil, fmt.Errorf("No result or duplicate patient %s", patient.Id())
		}
		delete(m, patient.Id())

		r, err := result.ToLabResult(patient)
		if err != nil {
			return nil, fmt.Errorf("Error creating lab result for patient %v %v", result, patient)
		}
		results[i] = r
	}

	if len(m) != 0 {
		keys := make([]string, len(m))
		for k := range m {
			keys = append(keys, k)
		}

		return nil, fmt.Errorf("Not all patients had results %v", keys)
	}

	return results, nil
}

func Load(patientsIn, resultsIn io.Reader) ([]*lab.Result, error) {

	patients, err := LoadPatients(patientsIn)
	if err != nil {
		return nil, err
	}

	rawResults, err := LoadLabResults(resultsIn)
	if err != nil {
		return nil, err
	}

	return HashJoin(rawResults, patients)
}

func LoadFile(patientsFileName, resultsFileName string) ([]*lab.Result, error) {
	patientsFile, err := os.Open(patientsFileName)
	if err != nil {
		return nil, err
	}
	defer patientsFile.Close()

	resultsFile, err := os.Open(resultsFileName)
	if err != nil {
		return nil, err
	}
	defer resultsFile.Close()

	return Load(patientsFile, resultsFile)
}
