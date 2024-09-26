package store

import (
	"database/sql"
	"errors"

	"github.com/desafio-ll/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) GetDentistByID(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE matricula = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.Nombre, &dentist.Apellido, &dentist.Matricula)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStore) GetPatientByID(id int) (domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Nombre, &patient.Apellido, &patient.Domicilio, &patient.Dni, &patient.FechaAlta)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStore) GetTurnByID(id int) (domain.Turn, error) {
	var turn domain.Turn
	query := "SELECT * FROM turns WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&turn.Id, &turn.PacienteID, &turn.DentistaID, &turn.FechaHora, &turn.Descripcion)
	if err != nil {
		return domain.Turn{}, err
	}
	return turn, nil
}

func (s *sqlStore) CreateDentist(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (nombre, apellido, matricula) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.Nombre, dentist.Apellido, dentist.Matricula)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) CreatePatient(patient domain.Patient) error {
	query := "INSERT INTO patients (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.Nombre, patient.Apellido, patient.Domicilio, patient.Dni, patient.FechaAlta)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) CreateTurn(turn domain.Turn) error {
	query := "INSERT INTO turns (paciente_id, dentista_id, fecha_hora, descripcion) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turn.PacienteID, turn.DentistaID, turn.FechaHora, turn.Descripcion)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) UpdateDentist(dentist domain.Dentist) error {
	query := "UPDATE dentists SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.Nombre, dentist.Apellido, dentist.Matricula, dentist.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) UpdatePatient(patient domain.Patient) error {
	query := "UPDATE patients SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.Nombre, patient.Apellido, patient.Domicilio, patient.Dni, patient.FechaAlta, patient.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) UpdateTurn(turn domain.Turn) error {
	query := "UPDATE turns SET patient_id = ?, dentist_id = ?, fecha_hora = ?, descripcion = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turn.PacienteID, turn.DentistaID, turn.FechaHora, turn.Descripcion, turn.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteDentist(id int) error {
	query := "DELETE FROM dentists WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeletePatient(id int) error {
	query := "DELETE FROM patients WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteTurn(id int) error {
	query := "DELETE FROM turns WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) ExistsDentist(matricula int) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentists WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

func (s *sqlStore) ExistsPatient(dni int) bool {
	var exists bool
	var id int
	query := "SELECT id FROM patients WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

func (s *sqlStore) GetDentistByMatricula(matricula int) (int, error) {
	var id int
	query := "SELECT id FROM dentists WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sqlStore) GetPatientByDni(dni int) (int, error) {
	var id int
	query := "SELECT id FROM patients WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sqlStore) CreateBodyTurn(bodyturn domain.BodyTurn) error {
	patient_id, err := s.GetPatientByDni(bodyturn.DniPaciente)
	if err != nil {
		return errors.New("patient not found")
	}
	dentist_id, err := s.GetDentistByMatricula(bodyturn.MatriculaDentista)
	if err != nil {
		return errors.New("dentist not found")
	}
	query := "INSERT INTO turns (paciente_id, dentista_id, fecha_hora, descripcion) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient_id, dentist_id, bodyturn.FechaHora, bodyturn.Descripcion)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetTurnByDniPatient(dni int) (domain.DetailTurn, error) {
	var detailturn domain.DetailTurn
	query := `SELECT t.fecha_hora, t.descripcion, 
	p.nombre, p.apellido, p.dni, p.fecha_alta, 
	d.nombre, d.apellido, d.matricula FROM turns t 
	INNER JOIN patients p ON t.paciente_id = p.id 
	INNER JOIN dentists d ON t.dentista_id = d.id WHERE p.dni = ?;`
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&detailturn.FechaHora, &detailturn.Descripcion,
		&detailturn.NombrePaciente, &detailturn.ApellidoPaciente, &detailturn.DniPaciente, &detailturn.FechaAltaPaciente,
		&detailturn.NombreDentista, &detailturn.ApellidoDentista, &detailturn.Matricula)
	if err != nil {
		return domain.DetailTurn{}, err
	}
	return detailturn, nil
}
