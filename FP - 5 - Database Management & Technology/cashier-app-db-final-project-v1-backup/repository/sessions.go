package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	err := u.db.Where("token = ?", tokenTarget).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	err := u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err = u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var result model.Session
	err := u.db.Model(&model.Session{}).Where("username = ?", name).Take(&result).Error
	if err != nil {
		return model.Session{}, err
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var result model.Session
	err := u.db.Model(&model.Session{}).Where("token = ?", token).Take(&result).Error
	if err != nil {
		return model.Session{}, err
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
