package sqlquere

var SyncTextTable = `WITH updated_rows AS (
	UPDATE text_data
	SET
	  name = $1,
	  data = $2,
	  uid = $3,
	  deleted = $4,
	  last_update = $5
	WHERE
	  name = $6
	  AND last_update < $7
	RETURNING *
  )
  INSERT INTO text_data (name, data, uid, deleted, last_update)
  SELECT
	$8,
	$9,
	$10,
	$11,
	$12
  WHERE
	NOT EXISTS (SELECT 1 FROM updated_rows)
	AND NOT EXISTS (SELECT *
	FROM text_data
	WHERE
	  name = $13
	  AND last_update >= $14);
  
  SELECT name, data, uid, deleted, last_update
	FROM text_data
	WHERE
	  name = $15
	  AND last_update >= $16`

var SyncAuthTable = `WITH updated_rows AS (
		UPDATE logins
		SET
		  name = $1,
		  login = $2,
		  password = $3,
		  uId = $4,
		  deleted = $5,
		  last_update = $6
		WHERE
		  name = $7
		  AND last_update < $8
		RETURNING *
	  )
	  INSERT INTO logins (name, login, password, uid, deleted, last_update)
	  SELECT
		$9,
		$10,
		$11,
		$12,
		$13,
		$14
	  WHERE
		NOT EXISTS (SELECT 1 FROM updated_rows)
		AND NOT EXISTS (SELECT *
		FROM logins
		WHERE
		  name = $15
		  AND last_update >= $16);
	  
	  SELECT name, login, password, uid, deleted, last_update
		FROM logins
		WHERE
		  name = $17
		  AND last_update >= $18`

var SyncBinTable = `WITH updated_rows AS (
			UPDATE binares_data
			SET
			  name = $1,
			  data = $2,
			  uId = $3,
			  deleted = $4,
			  last_update = $5
			WHERE
			  name = $6
			  AND last_update < $7
			RETURNING *
		  )
		  INSERT INTO binares_data (name, data, uid, deleted, last_update)
		  SELECT
			$8,
			$9,
			$10,
			$11,
			$12
		  WHERE
			NOT EXISTS (SELECT 1 FROM updated_rows)
			AND NOT EXISTS (SELECT *
			FROM binares_data
			WHERE
			  name = $13
			  AND last_update >= $14);
		  
		  SELECT name, data, uid, deleted, last_update
			FROM binares_data
			WHERE
			  name = $15
			  AND last_update >= $16`

var SyncCardTable = `WITH updated_rows AS (
				UPDATE cards
				SET
				  name = $1,
				  number = $2,
				  date = $3,
				  cvv = $4,
				  uId = $5,
				  deleted = $6,
				  last_update = $7
				WHERE
				  name = $8
				  AND last_update < $9
				RETURNING *
			  )
			  INSERT INTO cards (name, number, date, cvv, uid, deleted, last_update)
			  SELECT
				$10,
				$11,
				$12,
				$13,
				$14,
				$15,
				$16
			  WHERE
				NOT EXISTS (SELECT 1 FROM updated_rows)
				AND NOT EXISTS (SELECT *
				FROM cards
				WHERE
				  name = $17
				  AND last_update >= $18);
			  
			  SELECT name, number, date, cvv, uid, deleted, last_update
				FROM cards
				WHERE
				  name = $19
				  AND last_update >= $20`
