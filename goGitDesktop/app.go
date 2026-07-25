package main

import (
	"context"

	"github.com/AlexanderVallejos10/goforge/api"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCoreInfo() string {

	return "GoForge Core v0.1 conectado"

}

func (a *App) GetBranch(
	ruta string,
) string {

	return api.GetBranch(
		ruta,
	)

}

func (a *App) GetStatus(
	ruta string,
) interface{} {

	return api.GetStatus(
		ruta,
	)

}

func (a *App) Add(
	ruta string,
) string {

	return api.Add(
		ruta,
	)

}

func (a *App) Commit(
	ruta string,
	mensaje string,
) string {

	return api.Commit(
		ruta,
		mensaje,
	)

}
