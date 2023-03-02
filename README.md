# Go-Guide

Örnek olarak, "users" tablosuna yeni bir sütun eklemek için SQL sorgusu hazırlayalım ve daha sonra sorguyu çalıştıralım:

```
func addColumn(db *sql.DB, colName string, colType string) error {
    // SQL sorgusunu hazırla
    query := "ALTER TABLE users ADD COLUMN " + colName + " " + colType
    stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Sorguyu çalıştır
    _, err = stmt.Exec()
    if err != nil {
        return err
    }
    return nil
}
```
Yukarıdaki fonksiyon, "db.Prepare" fonksiyonunu kullanarak önceden hazırlanmış bir SQL sorgusu oluşturur ve "db.Exec" fonksiyonunu kullanarak sorguyu çalıştırır.

Ayrıca, "db.Exec" fonksiyonu kullanarak aynı işlemi tek bir satırda da gerçekleştirebiliriz:

```
func addColumn(db *sql.DB, colName string, colType string) error {
    // SQL sorgusunu hazırla ve çalıştır
    query := "ALTER TABLE users ADD COLUMN " + colName + " " + colType
    _, err := db.Exec(query)
    if err != nil {
        return err
    }
    return nil
}
```
Bu örnekte, "db.Exec" fonksiyonu kullanılarak sorgu doğrudan çalıştırılır. "db.Prepare" ve "db.Exec" fonksiyonları arasındaki temel fark, bir sorgunun kaç kez çalıştırılacağıdır. Eğer aynı sorgu birden fazla kez çalıştırılacaksa, "db.Prepare" fonksiyonu kullanılarak sorgu önceden hazırlanabilir ve daha sonra "db.Exec" fonksiyonu kullanılarak çalıştırılabilir. Bu, performansı artırabilir ve daha az kod yazmamızı sağlar.

Ayrıca başka fonksiyonlarda vardır.

"db.Query": SQL sorgusu çalıştırır ve sonuçları verir. Örneğin:

```
rows, err := db.Query("SELECT * FROM users")
```

"db.QueryRow": SQL sorgusu çalıştırır ve yalnızca bir satırın sonucunu verir. Örneğin:

```
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
```

"db.Exec": SQL sorgusu çalıştırır ve sonucu döndürmez. Örneğin:

```
result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "John", 1)
```

"tx.Exec": Bir işlem içinde (transaction) SQL sorgusu çalıştırır ve sonucu döndürmez. Örneğin:

```
tx, err := db.Begin()
result, err := tx.Exec("UPDATE users SET name = ? WHERE id = ?", "John", 1)
```

"stmt.Exec": Hazırlanan bir SQL ifadesini çalıştırır ve sonucu döndürmez. Örneğin:

```
stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
result, err := stmt.Exec("John", 30)
```

"stmt.Query": Hazırlanan bir SQL ifadesini çalıştırır ve sonuçları verir. Örneğin:

```
stmt, err := db.Prepare("SELECT * FROM users WHERE age > ?")
rows, err := stmt.Query(30)
```

Bu fonksiyonlar, Go dilinde "database/sql" paketinin kullanımı ile birlikte kullanılır ve SQL veritabanlarına erişmek için oldukça faydalıdır.
