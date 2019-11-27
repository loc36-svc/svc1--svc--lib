package lib

func Conf () (dbUser, dbPass string, connTimeout, writeTimeout, readTimeout uint16, dbmsPubKey string, err error) {
	/*  ALGORITHM

		step 100: load conf file
		step 110: if loading fails: handle error

		step 120: get data 'dbms_user_name'
		step 130: if no data is available: return error

		step 140: get data 'dbms_user_pass'
		step 150: if no data is available: return error

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
		step 260: if no data is available: return error
	*/

	// step 100 ..1.. {
	conf, errX := viperLib.NewFileViper (confFileName, "yaml")
	if errX != nil {
		err = err.New ("Unable to load conf file.", nil, nil, errX)
		return
	}
	// ..1.. }

	dbUser = conf.GetString ("dbms_user_name")

	// step 110 ..1.. {
	if dbUser == "" {
		err = err.New ("Conf data 'dbms_user_name': Data not provided.", nil, nil)
		return
	}
	// .. }

	dbPass = conf.GetString ("dbms_user_pass")

	// step 100 ..1.. {
	if dbPass == "" {
		err = err.New ("Conf data 'dbms_user_pass': Data not provided.", nil, nil)
		return
	}
	// .. }

	strConnTimeout := conf.GetString ("conn_timeout")

	timeoutF, errF := strconv.Atoi (conf.GetString ("conn_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'conn_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'conn_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'conn_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	connTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'wrte_timeout'.  ..1.. {
	if conf.GetString ("wrte_timeout") == "" {
		err = err.New ("Conf data 'wrte_timeout': Data not set.", nil, nil)
		return
	}
	timeoutF, errF := strconv.Atoi (conf.GetString ("wrte_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'wrte_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'wrte_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'wrte_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	writeTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'read_timeout'.  ..1.. {
	if conf.GetString ("read_timeout") == "" {
		err = err.New ("Conf data 'read_timeout': Data not set.", nil, nil)
		return
	}
	timeoutF, errF := strconv.Atoi (conf.GetString ("read_timeout"))
	if errF != nil {
		err = err.New ("Conf data 'read_timeout': Value seems invalid.", nil, nil, errF)
		return
	}
	if timeoutF == 0 {
		err = err.New ("Conf data 'read_timeout': Timeout can not be zero.", nil, nil)
		return
	}
	if timeoutF > 960 {
		err = err.New ("Conf data 'read_timeout': Timeout can not be greater than 960 seconds (16 minutes).", nil, nil)
		return
	}
	readTimeout = uint16 (timeoutF)
	// .. }

	// Processing conf data 'dbms_pub_key'.  ..1.. {	
	if conf.GetString ("dbms_pub_key") == "" {
		err = err.New ("Conf data 'dbms_pub_key': Data not set.", nil, nil)
		return
	}
	dbmsPubKey = conf.GetString ("dbms_pub_key")
	okJ, errJ := afero.Exists (afero.NewOsFs (), dbmsPubKey)
	if errJ != nil {
		err = err.New ("Conf data 'dbms_pub_key': Unable to confirm existence of file.", nil, nil, errJ)
		return
	}
	if okJ == false {
		err = err.New ("Conf data 'dbms_pub_key': File not found.", nil, nil)
		return
	}
	// .. }
}
var confFileName = "conf.yml"
