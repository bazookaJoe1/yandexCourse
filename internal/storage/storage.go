package storage

import (
	"errors"
	"fmt"
	"strconv"

	"log"
)

func (s *Storage) GetAll() string {
	s.Lock()
	defer s.Unlock()

	var outputString = ""

	for nameM, valueM := range s.values {
		var valueD uint = 0
		var valueS = ""

		valueD, ok := valueM.(uint)
		if !ok {
			valueD, _ := valueM.(float64)
			valueS = strconv.FormatFloat(valueD, 'f', 3, 64)
		} else {
			valueS = strconv.FormatUint(uint64(valueD), 10)
		}

		outputString += fmt.Sprintf("<BR>%s: %s</BR>", nameM, valueS)
	}

	return outputString

}

func (s *Storage) Get(TypeM, NameM string) (string, error) {
	s.Lock()
	defer s.Unlock()

	var err error = nil
	var ValueS = ""

	switch TypeM {
	case "counter":
		if ValueM, found := s.values[NameM].(uint); found {
			ValueS = strconv.FormatUint(uint64(ValueM), 10)
		} else {
			return "", fmt.Errorf("metric not found")
		}
	case "gauge":
		if ValueM, found := s.values[NameM].(float64); found {
			ValueS = strconv.FormatFloat(ValueM, 'f', 3, 64)
		} else {
			return "", fmt.Errorf("metric not found")
		}
	default:
		return "", fmt.Errorf("metric not found")
	}
	return ValueS, err
}

// Save(TypeM, NameM, ValueM string) according to TypeM (type of metric) performs check of ValueM (value of metric) and saves it to storage.
func (s *Storage) Save(TypeM, NameM, ValueM string) error {
	s.Lock()
	defer s.Unlock()

	var err error = nil

	switch TypeM {
	case "counter":
		if ValueU, err := checkCounter(ValueM); err != nil {
			err = errors.New("wrong type of counter")
			log.Printf("Wrong type of counter <%v>", ValueM)
			return err
		} else {
			if value, ok := s.values[NameM].(uint); ok {
				s.values[NameM] = ValueU + value
				log.Printf("Counter <%s> with value <%s> added.", NameM, ValueM)
			} else {
				s.values[NameM] = ValueU
				log.Printf("Counter <%s> with value <%s> saved.", NameM, ValueM)
			}

		}

	case "gauge":
		if ValueF, err := checkGauge(ValueM); err != nil {
			err = errors.New("wrong type of gauge")
			log.Printf("Wrong type of gauge <%v>", ValueM)
			return err
		} else {
			s.values[NameM] = ValueF
			log.Printf("Gauge metric <%s> with value <%s> saved.", NameM, ValueM)
		}
	default:
		err = errors.New("not implemented type of metric")
		log.Printf("Not implemented type of metric <%v>", TypeM)
		return err
	}

	return err
}

// StorageInit() returns a pointer to Storage struct that implements repositories interface.
func StorageInit() Repositories {
	return &Storage{values: make(map[string]any)}
}

// checkCounter(ValueM string) performs check of correct counter value in ValueM (value of metrics).
func checkCounter(ValueM string) (uint, error) {

	if ValueU, err := strconv.ParseUint(ValueM, 10, 64); err != nil {
		err := errors.New("wrong type of counter")
		log.Printf("Wrong type of counter <%v>", ValueM)
		return uint(ValueU), err
	} else {

		return uint(ValueU), err
	}
}

// checkGauge(ValueM string) performs check of correct gauge value in ValueM (value of metrics).
func checkGauge(ValueM string) (float64, error) {
	if ValueF, err := strconv.ParseFloat(ValueM, 64); err != nil {
		err := errors.New("wrong type of gauge")
		log.Printf("Wrong type of gauge <%v>", ValueM)
		return float64(ValueF), err
	} else {

		return float64(ValueF), err
	}
}
