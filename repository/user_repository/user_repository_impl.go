package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"go_share/helper"
	"go_share/model/domain"
	"strconv"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO ref_users (nama, email, password, id_lembaga, tiket) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, query, user.Nama, user.Email, user.Password, user.IdLembaga, user.Tiket)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.IdUser = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "UPDATE ref_users SET nama = ?, id_lembaga = ?, token = ?, no_hp = ?, jenjang_pendidikan = ?, bahasa = ?, alamat = ?, kompetensi = ? WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, query, user.Nama, user.IdLembaga, user.Token, user.NoHp, user.JenjangPendidikan, user.Bahasa, user.Alamat, user.Kompetensi, user.IdUser)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM ref_users WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	query := "SELECT id_user, id_socket, nama, email, password, id_lembaga, token, tiket, link_foto, no_hp, jenjang_pendidikan, alamat, bahasa, status, kompetensi, is_login FROM ref_users WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.IdSocket, &user.Nama, &user.Email, &user.Password, &user.IdLembaga, &user.Token, &user.Tiket, &user.LinkFoto, &user.NoHp, &user.JenjangPendidikan, &user.Alamat, &user.Bahasa, &user.Status, &user.Kompetensi, &user.IsLogin)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User dengan id " + strconv.Itoa(user.IdUser) + " tidak ada")
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT id_user, nama, email, password, link_foto, no_hp, jenjang_pendidikan, alamat, bahasa, kompetensi, is_login FROM ref_users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.Nama, &user.Email, &user.Password, &user.LinkFoto, &user.NoHp, &user.JenjangPendidikan, &user.Alamat, &user.Bahasa, &user.Kompetensi, &user.IsLogin)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User dengan email " + email + " tidak ada")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id_user, nama, email, link_foto, no_hp, jenjang_pendidikan, alamat, bahasa, kompetensi, is_login FROM ref_users"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.IdUser, &user.Nama, &user.Email, &user.LinkFoto, &user.NoHp, &user.JenjangPendidikan, &user.Alamat, &user.Bahasa, &user.Kompetensi, &user.IsLogin)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}
