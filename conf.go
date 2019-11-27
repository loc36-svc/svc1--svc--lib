package lib

import (
	"github.com/qamarian-dtp/err"
	"github.com/qamarian-lib/viper"
	"github.com/spf13/afero"
	"strconv"
)

func Conf () (dbmsUser, userPass string, connTimeout, wrteTimeout, readTimeout uint16, dbmsPubKey string, funcErr error) {
	/*  ALGORITHM

		step 100: load conf file
		step 110: if loading fails: handle error

		step 120: get data 'dbms_user_name'
		step 130: if no data was provided: return error

		step 140: get data 'dbms_user_pass'
		step 150: if no data was provided: return error

		step 160: get data 'conn_timeout'
		step 170: check data validity
		step 180: if data is invalid: return error

		step 190: get data 'wrte_timeout'
		step 200: check data validity
		step 210: if data is invalid: return error

		step 220: get data 'read_timeout'
		step 230: check data validity
		step 240: if data is invalid: return error

		step 250: get data 'dbms_pub_key'
		step 260: check if file exists
		step 270: if file does not exists: return error
	*/

	// step 100
	conf, errX := viper.NewFileViper (confFileName, "yaml")

	// step 110 .. {
	if errX != nil {
		funcErr = err.New ("Unable to load conf file.", nil, nil, errX)
		return
	}
	// .. }

	// step 120
	dbmsUser = conf.GetString ("dbms_user_name")

	// step 130 .. {
	if dbmsUser == "" {
		funcErr = err.New ("Conf data 'dbms_user_name': Data not provided.", nil, nil)
		return
	}
	// .. }

	// step 140
	userPass = conf.GetString ("dbms_user_pass")

	// step 150 .. {
	if userPass == "" {
		funcErr = err.New ("Conf data 'dbms_user_pass': Data not provided.", nil, nil)
		return
	}
	// .. }

	// step 160
	strConnTimeout := conf.GetString ("conn_timeout")

	// step 170 and 180 .. {
	intConnTimeout, errY := strconv.Atoi (strConnTimeout)
	if errY != nil || intConnTimeout == 0 || intConnTimeout > 960 {
		funcErr = err.New ("Conf data 'conn_timeout': Value seems invalid.", nil, nil, errY)
		return
	}
	connTimeout = uint16 (intConnTimeout)
	// .. }

	// step 190
	strWrteTimeout := conf.GetString ("wrte_timeout")

	// step 200 and 210 .. {
	intWrteTimeout, errA := strconv.Atoi (strWrteTimeout)
	if errA != nil || intWrteTimeout == 0 || intWrteTimeout > 960 {
		funcErr = err.New ("Conf data 'wrte_timeout': Value seems invalid.", nil, nil, errA)
		return
	}
	wrteTimeout = uint16 (intWrteTimeout)
	// .. }

	// step 220
	strReadTimeout := conf.GetString ("read_timeout")

	// step 230 and 240 .. {
	intReadTimeout, errB := strconv.Atoi (strReadTimeout)
	if errB != nil || intReadTimeout == 0 || intReadTimeout > 960 {
		funcErr = err.New ("Conf data 'read_timeout': Value seems invalid.", nil, nil, errB)
		return
	}
	readTimeout = uint16 (intReadTimeout)
	// .. }

	// step 250
	dbmsPubKey = conf.GetString ("dbms_pub_key")

	// step 260 .. {
	yes, errC := afero.Exists (afero.NewOsFs (), dbmsPubKey)
	if errC != nil {
		funcErr = err.New ("Conf data 'dbms_pub_key': Unable to confirm file's existence.", nil, nil, errC)
		return
	}
	// .. }

	// step 270 .. {
	if yes == false {
		funcErr = err.New ("Conf data 'dbms_pub_key': File does not exist.", nil, nil)
		return
	}
	// .. }

	return
}
var confFileName = "conf.yml"
