package geolocate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//UpdateDBs checks and downloads (if necessary) the latest City and ASN Geolite2 databases from the MaxMind site
//NB: Requires a MaxMind license key
func UpdateDBs(licenseKey string) error {

	if licenseKey == "" {
		return errors.New("geolocate: No license key provided")
	}
	//Update ASN
	_, err := geoDBUpdate(dbASNFileNameCompressed, fmt.Sprintf(dbASNMD5Path, licenseKey), fmt.Sprintf(dbASNDBPath, licenseKey))
	if err != nil {
		//log.Printf("ERROR: Could not complete update for %s: %s\n", dbASNFileNameCompressed, err)
		return err
	}

	//Update City
	_, err = geoDBUpdate(dbCityFileNameCompressed, fmt.Sprintf(dbCityMD5Path, licenseKey), fmt.Sprintf(dbCityDBPath, licenseKey))
	if err != nil {
		//log.Printf("ERROR: Could not complete update for %s: %s\n", dbCityFileNameCompressed, err)
		return err
	}
	return nil
}

func geoDBUpdate(filename string, md5path string, dbpath string) (int, error) {

	if fileExists(filename) {

		//Get MD5 of latest DB
		resp, err := http.Get(md5path)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		//Read body
		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}

		//Get hash of existing file
		dbMD5, err := fileHashMD5(filename)
		if err != nil {
			return 0, err
		}

		//Check MD5s
		if dbMD5 != string(out) {
			//log.Printf("Status: Update for %s available, downloading\n", filename)
			err = fileDownload(dbpath, filename)
			if err != nil {
				return 0, err
			}
		}

	} else { //Doesn't exist do download

		//log.Printf("Status: %s does not exist, downloading\n", filename)

		err := fileDownload(dbpath, filename)
		if err != nil {
			return 0, err
		}

	}

	return 1, nil

}
