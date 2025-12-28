package querymatch

import (
	"testing"
)

// TODO: generalizzare & trasformare in una libreria (modulo)
func TestMatch(t *testing.T) {
	nomeEstratto, err := Match("nome (.*) è valido", "nome mirko è valido")
	if err != nil {
		t.Errorf(err.Error())
	}
	if nomeEstratto[1] != "mirko" {
		t.Errorf("volevo \"mirko\" in el. 1, ottenuto \"%s\"", nomeEstratto[1])
	} else {
		t.Logf("ottenuto mirko OK")
	}
}
