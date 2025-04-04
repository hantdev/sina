package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hantdev/sina/bootstrap"
	api "github.com/hantdev/mitras/api/http"
	apiutil "github.com/hantdev/mitras/api/http/util"
	"github.com/hantdev/mitras/pkg/errors"
)

// EncodeError encodes an error response.
func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	var wrapper error
	if errors.Contains(err, apiutil.ErrValidation) {
		wrapper, err = errors.Unwrap(err)
	}

	w.Header().Set("Content-Type", api.ContentType)

	status, nerr := toStatus(err)
	if nerr != nil {
		err = unwrap(err)
		w.WriteHeader(status)
		encodeErrorMessage(err, wrapper, w)
		return
	}

	if wrapper != nil {
		err = errors.Wrap(wrapper, err)
	}
	api.EncodeError(ctx, err, w)
}

func toStatus(err error) (int, error) {
	switch {
	case errors.Contains(err, bootstrap.ErrExternalKey),
		errors.Contains(err, bootstrap.ErrExternalKeySecure):
		return http.StatusForbidden, err

	case errors.Contains(err, bootstrap.ErrBootstrapState),
		errors.Contains(err, bootstrap.ErrAddBootstrap):
		return http.StatusBadRequest, err

	case errors.Contains(err, bootstrap.ErrBootstrap):
		return http.StatusNotFound, err

	default:
		return 0, nil
	}
}

func encodeErrorMessage(err, wrapper error, w http.ResponseWriter) {
	if wrapper != nil {
		err = errors.Wrap(wrapper, err)
	}
	if errorVal, ok := err.(errors.Error); ok {
		if err := json.NewEncoder(w).Encode(errorVal); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func unwrap(err error) error {
	wrapper, err := errors.Unwrap(err)
	if wrapper != nil {
		return wrapper
	}
	return err
}