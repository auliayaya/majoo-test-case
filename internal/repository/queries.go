package repository

var (
	getMerchantDailyIncome = `
	WITH RECURSIVE all_dates(dt, m_id, merchant_name) AS (
		SELECT DATE(?)
			 , id AS m_id
			 , merchant_name
		  FROM Merchants
		 WHERE user_id = ?
		UNION ALL
		SELECT dt + INTERVAL 1 DAY, m_id, merchant_name
		  FROM all_dates
		 WHERE dt + INTERVAL 1 DAY <= ?
	)
	SELECT ad.m_id AS "merchant_id"
		 , ad.merchant_name AS "merchant_name"
		 , ad.dt AS "date"
		 , (SELECT IFNULL(SUM(t.bill_total), 0)
			  FROM Transactions AS t
			 WHERE t.merchant_id = ad.m_id
			   AND DATE(t.created_at) = ad.dt
		   ) AS "total"
	  FROM all_dates ad
	  LIMIT ? OFFSET ?
	`
	getMerchanDailyIncomeNoLimitOffset = `
	WITH RECURSIVE all_dates(dt, m_id, merchant_name) AS (
		SELECT DATE(?)
			 , id AS m_id
			 , merchant_name
		  FROM Merchants
		 WHERE user_id = ?
		UNION ALL
		SELECT dt + INTERVAL 1 DAY, m_id, merchant_name
		  FROM all_dates
		 WHERE dt + INTERVAL 1 DAY <= ?
	)
	SELECT ad.m_id AS "merchant_id"
		 , ad.merchant_name AS "merchant_name"
		 , ad.dt AS "date"
		 , (SELECT IFNULL(SUM(t.bill_total), 0)
			  FROM Transactions AS t
			 WHERE t.merchant_id = ad.m_id
			   AND DATE(t.created_at) = ad.dt
		   ) AS "total"
	  FROM all_dates ad
	`
	getOutletDailyIncome = `
	WITH RECURSIVE all_dates(dt, merchant_id, merchant_name, outlet_id, outlet_name) AS (
		SELECT DATE(?)
			 , m.id AS merchant_id
			 , m.merchant_name
			 , o.id AS outlet_id
			 , o.outlet_name
		  FROM Outlets o
		  LEFT JOIN Merchants m
			ON m.id = o.merchant_id
		 WHERE m.user_id = ?
		
		UNION ALL
		
		SELECT dt + INTERVAL 1 DAY
			 , merchant_id
			 , merchant_name
			 , outlet_id
			 , outlet_name
		  FROM all_dates
		 WHERE dt + INTERVAL 1 DAY <= ?
	)
	SELECT ad.merchant_id
		 , ad.merchant_name
		 , ad.outlet_id
		 , ad.outlet_name
		 , ad.dt AS "date"
		 , (SELECT IFNULL(SUM(t.bill_total), 0)
			  FROM Transactions AS t
			 WHERE t.merchant_id = ad.merchant_id
			   AND t.outlet_id = ad.outlet_id
			   AND DATE(t.created_at) = ad.dt
		   ) AS "total"
	  FROM all_dates ad
	  LIMIT ? OFFSET ?
	`
	getOutletDailyIncomeNoLimitOffset = `
	WITH RECURSIVE all_dates(dt, merchant_id, merchant_name, outlet_id, outlet_name) AS (
		SELECT DATE(?)
			 , m.id AS merchant_id
			 , m.merchant_name
			 , o.id AS outlet_id
			 , o.outlet_name
		  FROM Outlets o
		  LEFT JOIN Merchants m
			ON m.id = o.merchant_id
		 WHERE m.user_id = ?
		
		UNION ALL
		
		SELECT dt + INTERVAL 1 DAY
			 , merchant_id
			 , merchant_name
			 , outlet_id
			 , outlet_name
		  FROM all_dates
		 WHERE dt + INTERVAL 1 DAY <= ?
	)
	SELECT ad.merchant_id
		 , ad.merchant_name
		 , ad.outlet_id
		 , ad.outlet_name
		 , ad.dt AS "date"
		 , (SELECT IFNULL(SUM(t.bill_total), 0)
			  FROM Transactions AS t
			 WHERE t.merchant_id = ad.merchant_id
			   AND t.outlet_id = ad.outlet_id
			   AND DATE(t.created_at) = ad.dt
		   ) AS "total"
	  FROM all_dates ad
	`
	getOutletByIDAndUserID = `SELECT Outlets.* FROM Outlets LEFT JOIN Merchants ON Outlets.merchant_id = Merchants.id WHERE Outlets.id = ?  AND Merchants.user_id = ?`
)
