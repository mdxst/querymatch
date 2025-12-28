package querymatch

import (
	"regexp"
	"hash/adler32"
)

// map[checksum stringa controllo regexp]ogg. regexp
var cacheRegexp map[uint32]*regexp.Regexp

// data una stringa m del tipo: "ciao (.*)!" e una stringa s del tipo "ciao
// mirko!", Match ritornerà ["mirko"], nil. TIENE CACHE DEI REGEXP COMPILATI!
//
// COME LEGGGERE L'ARRAY RITORNATO:
// - IL PRIMO elemento è sempre la stringa "s" (el. 0)
// - DAL SECONDO IN POI ci sono le stringhe che hanno combaciato i punti con
//	"(.*)" o altri jolly regexp (el. 1>)
func Match(m, s string) ([]string, error) {
	val, ok := cacheRegexp[checksumChiave(m)]
	var r *regexp.Regexp
	if ok {
		// abbiamo già sto regexp compilato in cache, usiamolo...
		r = val
	} else {
		// il regexp va compilato!
		var err error
		r, err = PreparaRegexp(m, true)
		if err != nil {
			return nil, err
		}
	}

	return r.FindStringSubmatch(s), nil
}

// data una stringa c, ne ritorna la checksum "Adler-32" --> perchè "Adler-32"?
// È l'algoritmo più veloce, ma non è completamente affidabile....
// (compromesso)
func checksumChiave(c string) uint32 {
	return adler32.Checksum([]byte(c))
}

// data una stringa m, la compila regexp e ne ritorna l'oggetto pronto. In caso
// di errore, ritorna nil & un errore in 2° valore
//
// PARAMETRI:
// @param m: stringa da compilare regexp
// @param salvaCache: 'true' se vuoi che la salviamo nella nostra cache
// interna, cacheRegexp, per uso futuro
func PreparaRegexp(m string, salvaCache bool) (*regexp.Regexp, error) {
	r, err := regexp.Compile(m)
	if err != nil {
		return nil, err
	}
	return r, nil
}
