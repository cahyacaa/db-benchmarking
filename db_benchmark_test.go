package main

import "testing"

func BenchmarkPostgresQuery(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := postgresDB.Query("SELECT * FROM sales WHERE product_id = 1")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMariaDBQuery(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := mariaDB.Query("SELECT * FROM sales WHERE product_id = 1")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPostgresComplexQuery(b *testing.B) {
	query := `
        SELECT 
            c.customer_id,
            c.first_name,
            c.last_name,
            COUNT(o.order_id) AS total_orders,
            SUM(oi.price * oi.quantity) AS total_sales,
            AVG(oi.price * oi.quantity) AS average_order_value
        FROM customers c
        JOIN orders o ON c.customer_id = o.customer_id
        JOIN order_items oi ON o.order_id = oi.order_id
        WHERE o.order_date BETWEEN '2022-01-01' AND '2022-12-31'
        GROUP BY c.customer_id
        ORDER BY total_sales DESC;
    `
	for n := 0; n < b.N; n++ {
		_, err := postgresDB.Query(query)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMariaDBComplexQuery(b *testing.B) {
	query := `
        SELECT 
            c.customer_id,
            c.first_name,
            c.last_name,
            COUNT(o.order_id) AS total_orders,
            SUM(oi.price * oi.quantity) AS total_sales,
            AVG(oi.price * oi.quantity) AS average_order_value
        FROM customers c
        JOIN orders o ON c.customer_id = o.customer_id
        JOIN order_items oi ON o.order_id = oi.order_id
        WHERE o.order_date BETWEEN '2022-01-01' AND '2022-12-31'
        GROUP BY c.customer_id
        ORDER BY total_sales DESC;
    `
	for n := 0; n < b.N; n++ {
		_, err := mariaDB.Query(query)
		if err != nil {
			b.Fatal(err)
		}
	}
}
