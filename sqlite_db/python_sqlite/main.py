import sqlite3

def main():
    # Conectar a la base de datos SQLite
    conn = sqlite3.connect('users.db')
    cursor = conn.cursor()

    # Realizar algunas consultas
    print("Consultando algunos datos de usuarios:")

    # Consulta 1: Obtener todos los usuarios
    print("\nTodos los usuarios:")
    cursor.execute("SELECT * FROM users LIMIT 1000000")
    rows = cursor.fetchall()
    for row in rows:
        print(row)

    # Consulta 2: Obtener usuarios mayores de 30 años
    print("\nUsuarios mayores de 30 años:")
    cursor.execute("SELECT * FROM users WHERE Age > ? LIMIT 1000000", (30,))
    rows = cursor.fetchall()
    for row in rows:
        print(row)

    # Cerrar la conexión con la base de datos
    conn.close()

if __name__ == "__main__":
    main()
