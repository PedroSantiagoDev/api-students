# api-students
API to manage 'Golang do Zero' course students

Routes: 
- GET /students - List all students
- POST /students - Create students
- GET /students/:id - Get infos from a specific students
- PUT /students/:id - Update students
- DELETE /students/:id - Delete students

Struct Students:
- Name   string
- CPF    string
- Email  string
- Age    int
- Active string