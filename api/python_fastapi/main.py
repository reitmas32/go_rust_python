from fastapi import FastAPI
import sqlite3

app = FastAPI()

# Función para obtener un usuario por su ID desde la base de datos SQLite
def get_user_by_id(user_id: int):
    # Conectar a la base de datos SQLite
    conn = sqlite3.connect('users.db')
    cursor = conn.cursor()
    
    # Ejecutar la consulta para obtener el usuario por su ID
    cursor.execute("SELECT * FROM users WHERE id = ?", (user_id,))
    user = cursor.fetchone()
    
    # Cerrar la conexión con la base de datos
    conn.close()
    
    return user

@app.get("/users/{user_id}")
async def get_user(user_id: int):
    # Obtener el usuario por su ID desde la base de datos
    user = get_user_by_id(user_id)
    
    if user:
        # Si se encontró el usuario, devolver sus datos
        return {"id": user[0], "name": user[1], "age": user[2]}
    else:
        # Si no se encontró el usuario, devolver un mensaje de error
        return {"error": "User not found"}
