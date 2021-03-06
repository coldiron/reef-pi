package doser

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/reef-pi/reef-pi/controller/utils"
)

func (c *Controller) LoadAPI(r *mux.Router) {

	// swagger:route GET /api/doser/pumps Doser doserList
	// List all dosers.
	// List all dosers in reef-pi.
	// responses:
	// 	200: body:[]pump
	r.HandleFunc("/api/doser/pumps", c.list).Methods("GET")

	// swagger:operation GET /api/doser/pumps/{id} Doser doserGet
	// Get a doser by id.
	// Get an existing doser.
	// ---
	// parameters:
	//  - in: path
	//    name: id
	//    description: The Id of the doser
	//    required: true
	//    schema:
	//     type: integer
	// responses:
	//  200:
	//   description: OK
	//   schema:
	//    $ref: '#/definitions/pump'
	//  404:
	//   description: Not Found
	r.HandleFunc("/api/doser/pumps/{id}", c.get).Methods("GET")

	// swagger:operation PUT /api/doser/pumps Doser doserCreate
	// Create a doser.
	// Create a new doser.
	// ---
	// parameters:
	//  - in: body
	//    name: doser
	//    description: The doser to create
	//    required: true
	//    schema:
	//     $ref: '#/definitions/pump'
	// responses:
	//  200:
	//   description: OK
	r.HandleFunc("/api/doser/pumps", c.create).Methods("PUT")

	// swagger:operation POST /api/doser/pumps/{id} Doser doserUpdate
	// Update a doser.
	// Update an existing doser.
	//---
	//parameters:
	// - in: path
	//   name: id
	//   description: The Id of the doser to update
	//   required: true
	//   schema:
	//    type: integer
	// - in: body
	//   name: doser
	//   description: The doser to update
	//   required: true
	//   schema:
	//    $ref: '#/definitions/pump'
	//responses:
	// 200:
	//  description: OK
	// 404:
	//  description: Not Found
	r.HandleFunc("/api/doser/pumps/{id}", c.update).Methods("POST")

	// swagger:operation DELETE /api/doser/pumps/{id} Doser doserDelete
	// Delete a doser.
	// Delete an existing doser.
	// ---
	// parameters:
	//  - in: path
	//    name: id
	//    description: The Id of the doser to delete
	//    required: true
	//    schema:
	//     type: integer
	// responses:
	//  200:
	//   description: OK
	r.HandleFunc("/api/doser/pumps/{id}", c.delete).Methods("DELETE")

	// swagger:operation GET /api/doser/pumps/{id}/usage Doser doserUsage
	// Get usage history.
	// Get usage history for a given Doser.
	// ---
	// parameters:
	//  - in: path
	//    name: id
	//    description: The Id of the doser
	//    required: true
	//    schema:
	//     type: integer
	// responses:
	//  200:
	//   description: OK
	//  404:
	//   description: Not Found
	r.HandleFunc("/api/doser/pumps/{id}/usage", c.getUsage).Methods("GET")

	// swagger:operation POST /api/doser/pumps/{id}/calibrate Doser doserCalibrate
	// Calibrate a doser.
	// Calibrate a doser.
	// ---
	// parameters:
	//  - in: path
	//    name: id
	//    description: The Id of the doser
	//    required: true
	//    schema:
	//     type: integer
	//  - in: body
	//    name:
	//    description:
	//    required: true
	//    schema:
	//     $ref: '#/definitions/doserCalibrationDetails'
	// responses:
	//  200:
	//   description: OK
	//  404:
	//   description: Not Found
	r.HandleFunc("/api/doser/pumps/{id}/calibrate", c.calibrate).Methods("POST")

	// swagger:operation POST /api/doser/pumps/{id}/schedule Doser doserSchedule
	// Schedule dosing.
	// Schedule dosing.
	// ---
	// parameters:
	//  - in: path
	//    name: id
	//    description: The Id of the doser
	//    required: true
	//    schema:
	//     type: integer
	//  - in: body
	//    name:
	//    description:
	//    required: true
	//    schema:
	//     $ref: '#/definitions/dosingRegiment'
	// responses:
	//  200:
	//   description: OK
	//  404:
	//   description: Not Found
	r.HandleFunc("/api/doser/pumps/{id}/schedule", c.schedule).Methods("POST")
}

func (c *Controller) list(w http.ResponseWriter, r *http.Request) {
	fn := func() (interface{}, error) {
		return c.List()
	}
	utils.JSONListResponse(fn, w, r)
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	var p Pump
	fn := func() error {
		return c.Create(p)
	}
	utils.JSONCreateResponse(&p, fn, w, r)
}

func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	fn := func(id string) (interface{}, error) {
		return c.Get(id)
	}
	utils.JSONGetResponse(fn, w, r)
}

func (c *Controller) delete(w http.ResponseWriter, r *http.Request) {
	fn := func(id string) error {
		return c.Delete(id)
	}
	utils.JSONDeleteResponse(fn, w, r)
}

func (c *Controller) calibrate(w http.ResponseWriter, r *http.Request) {
	var cal CalibrationDetails
	fn := func(id string) error {
		return c.Calibrate(id, cal)
	}
	utils.JSONUpdateResponse(&cal, fn, w, r)
}

func (c *Controller) update(w http.ResponseWriter, r *http.Request) {
	var p Pump
	fn := func(id string) error {
		return c.Update(id, p)
	}
	utils.JSONUpdateResponse(&p, fn, w, r)
}

func (c *Controller) schedule(w http.ResponseWriter, r *http.Request) {
	var reg DosingRegiment
	fn := func(id string) error {
		return c.Schedule(id, reg)
	}
	utils.JSONUpdateResponse(&reg, fn, w, r)
}

func (c *Controller) getUsage(w http.ResponseWriter, req *http.Request) {
	fn := func(id string) (interface{}, error) { return c.statsMgr.Get(id) }
	utils.JSONGetResponse(fn, w, req)
}
