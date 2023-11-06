package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

func NewHandler(dinoService DinoService, logger *zerolog.Logger) http.Handler {

	router := chi.NewRouter()
	router.Route("/v1/", func(r chi.Router) {
		r.Get("/dinosaurs", getDinosHttp(dinoService, logger))
		r.Get("/dinosaurs/cage/{cageId}", getDinosByCageHttp(dinoService, logger))
		r.Get("/dinosaur/{dinoId}", getDinoHttp(dinoService, logger))
		r.Post("/dinosaur", addDinoHttp(dinoService, logger))
		r.Put("/dinosaur/{dinoId}", updateDinoHttp(dinoService, logger))
		r.Get("/cages", getCagesHttp(dinoService, logger))
		r.Get("/cage/{cageId}", getCageHttp(dinoService, logger))
		r.Post("/cage", addCageHttp(dinoService, logger))
		r.Put("/cage/{cageId}", updateCageHttp(dinoService, logger))
	})
	return router
}

// getDinosHttp gets all dinosaurs and returns result as json
func getDinosHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dinos, err := dinoService.GetDinos(r.Context())
		if err != nil {
			logger.Error().Err(err).Msg("error getting dino")
			err = render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		err = respondwithJSON(w, http.StatusOK, &dinos)
		if err != nil {
			err := render.Render(w, r, ServerError(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// getDinosByCageHttp gets all dinosaurs by cageId and returns result as json
func getDinosByCageHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cageId, _ := url.PathUnescape(chi.URLParam(r, "cageId"))
		id, err := strconv.ParseInt(cageId, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("error getting cageId")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		dinos, err := dinoService.GetDinosByCage(r.Context(), id)
		if err != nil {
			logger.Error().Err(err).Msg("error getting dinos by cage")
			if errors.Is(err, sql.ErrNoRows) {
				err := render.Render(w, r, NotFound(errors.New("not found")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
				return
			}
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		err = respondwithJSON(w, http.StatusOK, &dinos)
		if err != nil {
			logger.Error().Err(err).Msg("error getting dinos by cage")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// getDinoHttp gets all dinosaurs by cageId and returns result as json
func getDinoHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dinoId, _ := url.PathUnescape(chi.URLParam(r, "dinoId"))
		id, err := strconv.ParseInt(dinoId, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("error getting dino")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		dino, err := dinoService.GetDinoById(r.Context(), id)
		if err != nil {
			logger.Error().Err(err).Msg("error getting dino")
			if errors.Is(err, sql.ErrNoRows) {
				err := render.Render(w, r, NotFound(errors.New("not found")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
				return
			}
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		err = respondwithJSON(w, http.StatusOK, &dino)
		if err != nil {
			logger.Error().Err(err).Msg("error getting dino")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// getCageHttp gets cage by cageId and returns result as json
func getCageHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cageId, _ := url.PathUnescape(chi.URLParam(r, "cageId"))
		id, err := strconv.ParseInt(cageId, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("error parsing cageId")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		cage, err := dinoService.GetCageById(r.Context(), id)
		if err != nil {
			logger.Error().Err(err).Msg("error getting cage")
			if errors.Is(err, sql.ErrNoRows) {
				err := render.Render(w, r, NotFound(errors.New("not found")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
				return
			}
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		err = respondwithJSON(w, http.StatusOK, &cage)
		if err != nil {
			logger.Error().Err(err).Msg("error getting cage")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// addDinoHttp adds a new dino to the db
func addDinoHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		dino := Dinosaur{}
		err = json.Unmarshal(body, &dino)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data into dino struct")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}

		err = dinoService.AddDino(ctx, dino)
		if err != nil {
			logger.Error().Err(err).Msg("error saving dino")
			var serviceErr *ServiceRequestError
			if errors.As(err, &serviceErr) {
				err := render.Render(w, r, BadRequest(errors.New(err.(*ServiceRequestError).response)))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			} else {
				err := render.Render(w, r, ServerError(errors.New("server error")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			}
			return
		}

		err = respondwithJSON(w, http.StatusCreated, "{}")
		if err != nil {
			err := render.Render(w, r, ServerError(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// addCageHttp adds a new cage to the db
func addCageHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		cage := Cage{}
		err = json.Unmarshal(body, &cage)

		if err != nil {
			logger.Error().Err(err).Msg("error reading request data into cage struct")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}

		err = dinoService.AddCage(ctx, cage)
		if err != nil {
			logger.Error().Err(err).Msg("error saving cage")
			var serviceErr *ServiceRequestError
			if errors.As(err, &serviceErr) {
				err := render.Render(w, r, BadRequest(errors.New(err.(*ServiceRequestError).response)))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			} else {
				err := render.Render(w, r, ServerError(errors.New("server error")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			}
			return
		}

		err = respondwithJSON(w, http.StatusCreated, "{}")
		if err != nil {
			err := render.Render(w, r, ServerError(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// getCagesHttp gets all cages and returns result as json
func getCagesHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cages, err := dinoService.GetCages(r.Context())
		if err != nil {
			logger.Error().Err(err).Msg("error getting cages")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		err = respondwithJSON(w, 200, &cages)
		if err != nil {
			logger.Error().Err(err).Msg("error getting cages")
			err := render.Render(w, r, ServerError(errors.New("server error")))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// updateCageHttp updates a cage by cageId
func updateCageHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cageId, _ := url.PathUnescape(chi.URLParam(r, "cageId"))
		id, err := strconv.ParseInt(cageId, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("error parsing cageId")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}

		ctx := r.Context()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		cage := Cage{}
		err = json.Unmarshal(body, &cage)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data into cage struct")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		cage.Id = id

		err = dinoService.UpdateCage(ctx, cage)
		if err != nil {
			logger.Error().Err(err).Msg("error updating cage")
			var serviceErr *ServiceRequestError
			if errors.As(err, &serviceErr) {
				err := render.Render(w, r, BadRequest(errors.New(err.(*ServiceRequestError).response)))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			} else {
				err := render.Render(w, r, ServerError(errors.New("server error")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			}
			return
		}

		err = respondwithJSON(w, http.StatusOK, "{}")
		if err != nil {
			err := render.Render(w, r, ServerError(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

// updateDinoHttp updates a dino by id
func updateDinoHttp(dinoService DinoService, logger *zerolog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dinoId, _ := url.PathUnescape(chi.URLParam(r, "dinoId"))
		id, err := strconv.ParseInt(dinoId, 10, 64)
		if err != nil {
			logger.Error().Err(err).Msg("error parsing dinoId")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}

		ctx := r.Context()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		dino := Dinosaur{}
		err = json.Unmarshal(body, &dino)
		if err != nil {
			logger.Error().Err(err).Msg("error reading request data into dino struct")
			err := render.Render(w, r, BadRequest(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
		dino.Id = id
		err = dinoService.UpdateDino(ctx, dino)
		if err != nil {
			logger.Error().Err(err).Msg("error updating dino")
			var serviceErr *ServiceRequestError
			if errors.As(err, &serviceErr) {
				err := render.Render(w, r, BadRequest(errors.New(err.(*ServiceRequestError).response)))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			} else {
				err := render.Render(w, r, ServerError(errors.New("server error")))
				if err != nil {
					logger.Error().Err(err).Msg("render error")
				}
			}
			return
		}

		err = respondwithJSON(w, http.StatusOK, "{}")
		if err != nil {
			err := render.Render(w, r, ServerError(err))
			if err != nil {
				logger.Error().Err(err).Msg("render error")
			}
			return
		}
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	return err
}
