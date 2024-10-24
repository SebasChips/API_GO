package main

import (
	"fmt"

	"github.com/SebasChips/proyecto_final_go/database"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type estudiantes struct {
	Id    int16  `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:60;not null"`
	Group string `gorm:"size:5;not null"`
	Email string `gorm:"size:100;not null"`
}

type materias struct {
	Id   int16  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:30;not null"`
}

type calificaciones struct {
    Id         int16    `gorm:"primaryKey;autoIncrement"`
    StudentId  int16    `gorm:"column:studentId;not null"`
    Grade      float32 `gorm:"not null"`
    SubjectsId int16    `gorm:"column:subjectsId;not null"`
}

func main() {
	//API
	router := gin.Default()

	   // Configuración de CORS
	   corsConfig := cors.Config{
        AllowAllOrigins: true, // Permite todas las solicitudes de origen
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
    }
    router.Use(cors.New(corsConfig))

	fmt.Println("Server corriendo...")

	//Conexión base de datos
	db, err := database.NewDataBaseDriver()
	if err != nil {
		fmt.Println("Error al conectar a la BD en main", err)
		return
	}

	// Migrar la estructura a la base de datos
	db.AutoMigrate(&estudiantes{})
	db.AutoMigrate(&materias{})
	db.AutoMigrate(&calificaciones{})
	// Cargar plantillas HTML desde la carpeta "templates"
		router.LoadHTMLGlob("templates/*")
		router.Static("/static", "./static")
	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	router.GET("/alumnos.html", func(c *gin.Context) {
		c.HTML(200, "alumnos.html", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/calificaciones.html", func(c *gin.Context) {
        c.HTML(200, "calificaciones.html", gin.H{
            "title": "Alumnos",
        })
    })

	router.GET("/materias.html", func(c *gin.Context) {
		c.HTML(200, "materias.html", gin.H{
			"title": "Home Page",
		})
	})

	//Rutas estudiantes
	router.POST("/api/students", func(c *gin.Context) {

		var alumno estudiantes

		// Capturar datos JSON del cuerpo de la solicitud
		if err := c.BindJSON(&alumno); err != nil {
			c.JSON(400, gin.H{"error": "Error al procesar la solicitud"})
			return
		}
		
			result := db.Create(&alumno)
	
			if result.Error != nil {
				fmt.Println("Error al insertar datos:", result.Error)
			} else {
				c.JSON(200, alumno)
			}
	
	

	})

	router.DELETE("/api/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		var studentID int16

		fmt.Sscan(id, &studentID)

		student := estudiantes{Id: studentID}

		result := db.Delete(&student)

		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "Estudiante no encontrado"})
			return
		} else {
			c.JSON(200, gin.H{"message": "Estudiante eliminado"})
		}

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	})

	router.PUT("/api/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		var studentId int16
	
		if _, err := fmt.Sscan(id, &studentId); err != nil {
			c.JSON(400, gin.H{"error": "ID inválido"})
			return
		}
	
		// Buscar materia por ID
		var student estudiantes
		if err := db.First(&student, studentId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Materia no encontrada"})
			return
		}
	
		// Recuperar objeto con el JSON
		var updateData estudiantes
		if err := c.BindJSON(&updateData); err != nil {
			c.JSON(400, gin.H{"error": "Error al procesar la solicitud"})
			return
		}
		
		student.Group = updateData.Group
		student.Email = updateData.Email

		
	
		// Guardar los cambios en la base de datos
		if result := db.Save(&student); result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	
		c.JSON(200, gin.H{"message": "Alumno actualizado correctamente"})
	})

	router.GET("/api/students/", func(c *gin.Context) {
		var students []struct {
			Id int16
			Name string
			Group string
			Email string
		}

		db.Table("estudiantes").Select("`id`, `name`, `group`, `email`").Scan(&students)

		if students != nil {
			c.JSON(200, students)
		} else {
			c.JSON(404, "No se encontraron estudiantes")
		}

	})

	router.GET("/api/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		var studentID int16

		// Convertir id a int16
		fmt.Sscan(id, &studentID)
		var student estudiantes
		if err := db.First(&student, studentID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Estudiante no encontrado"})
			return
		}

		// Consultar la base de datos para obtener todos los registros
		result := db.Find(&student)

		if result != nil {
			c.JSON(200, student)
		} else {
			c.JSON(404, "No se encontro el estudiante")
		}

	})

	//////////////////////////////////////
	//RUTAS materias
	router.POST("/api/materias", func(c *gin.Context) {
		var subject materias

    // Capturar datos JSON del cuerpo de la solicitud
    if err := c.BindJSON(&subject); err != nil {
        c.JSON(400, gin.H{"error": "Error al procesar la solicitud"})
        return
    }
		

		result := db.Create(&subject)

		if result.Error != nil {
			fmt.Println("Error al insertar datos:", result.Error)
		} else {
			c.JSON(200, subject)
		}

	})

	router.GET("/api/materias/:id", func(c *gin.Context) {
		id := c.Param("id")

		var subjectID int16
		// Convertir id a int16
		fmt.Sscan(id, &subjectID)

		var materia struct {
			Name string
		}

		if err := db.Table("materias").Select("id, name").Where("id = ?", subjectID).Scan(&materia).Error; err != nil {
			c.JSON(404,gin.H{"error": "Materia no encontrada"})
			return
		}

		c.JSON(200, materia)
	})

	router.DELETE("/api/materias/:id", func(c *gin.Context) {
		id := c.Param("id")

		var subjectID int16
		// Convertir id a int16
		fmt.Sscan(id, &subjectID)

		// Crear una instancia del estudiante con el ID
		subject := materias{Id: subjectID}

		// Buscar y eliminar el registro
		result := db.Delete(&subject)

		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "Materia no encontrada"})
			return
		} else {
			c.JSON(200, gin.H{"message": "Materia eliminada"})
		}

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	})

	router.PUT("/api/materias/:id", func(c *gin.Context) {
		id := c.Param("id")
		var SubjectId int16
	
		// Convertir id a int16 de manera más robusta
		if _, err := fmt.Sscan(id, &SubjectId); err != nil {
			c.JSON(400, gin.H{"error": "ID inválido"})
			return
		}
	
		// Buscar materia por ID
		var subject materias
		if err := db.First(&subject, SubjectId).Error; err != nil {
			c.JSON(404, gin.H{"error": "Materia no encontrada"})
			return
		}
	
		// Recuperar objeto con el JSON
		var updateData materias
		if err := c.BindJSON(&updateData); err != nil {
			c.JSON(400, gin.H{"error": "Error al procesar la solicitud"})
			return
		}
		
		subject.Name = updateData.Name
		
	
		// Guardar los cambios en la base de datos
		if result := db.Save(&subject); result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	
		c.JSON(200, gin.H{"message": "Materia actualizada correctamente"})
	})
	
//////////////////////////////////////
	//RUTAS calificaciones
	router.POST("/api/grades", func(c *gin.Context) {
		var grade calificaciones
	
		// Capturar datos JSON del cuerpo de la solicitud
		if err := c.BindJSON(&grade); err != nil {
			c.JSON(400, gin.H{"error": "Error al procesar la solicitud"})
			return
		}

	
		result := db.Create(&grade)
	
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Error al insertar datos en la base de datos"})
			return
		}
	
		c.JSON(200, grade)
	})
	

	router.PUT("/api/grades/:grade_id", func(c *gin.Context) {
		id := c.Param("grade_id")
		var GradeId int16

		// Convertir id a int16
		fmt.Sscan(id, &GradeId)

		var gradeU calificaciones
		if err := db.First(&gradeU, GradeId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Calificación no encontrada"})
		return
	}


	var nuevaCalificacion struct {
		Grade float32 `json:"grade"`  
	}

		if err := c.BindJSON(&nuevaCalificacion); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}

		gradeU.Grade = nuevaCalificacion.Grade
		result := db.Save(&gradeU)

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Calificación actualizada"})
	})

	router.DELETE("/api/grades/:grade_id", func(c *gin.Context) {
		id := c.Param("grade_id")

		var gradeId int16
		// Convertir id a int16
		fmt.Sscan(id, &gradeId)

		// Crear una instancia del estudiante con el ID
		grade := calificaciones{Id: gradeId}

		// Buscar y eliminar el registro
		result := db.Delete(&grade)

		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "Calificación no encontrada"})
			return
		} else {
			c.JSON(200, gin.H{"message": "Calificación eliminada"})
		}

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	})

	router.GET("/api/grades/:grade_id/student/:student_id", func(c *gin.Context) {
		idGrade := c.Param("grade_id")
		idStudent := c.Param("student_id")
	
		var gradeID int16
		var studentID int16
	
		fmt.Sscan(idGrade, &gradeID)
		fmt.Sscan(idStudent, &studentID)
	
		var calificacion calificaciones
	
		if err := db.Table("calificaciones").
			Select("studentId, grade, subjectsId").Where("id = ? AND studentId = ?", gradeID, studentID).Scan(&calificacion).Error; err != nil {
			c.JSON(404, gin.H{"error": "Materia no encontrada"})
			return
		}
	
		fmt.Printf("Resultados de la consulta: %+v\n", calificacion)
	
		if calificacion.Grade == 0 && calificacion.StudentId == 0 && calificacion.SubjectsId == 0 {
			c.JSON(404, gin.H{"error": "Materia no encontrada o datos incorrectos"})
		} else {
			c.JSON(200, calificacion)
		}
	})
	

	router.GET("/api/grades/student/:student_id", func(c *gin.Context) {
		idStudent := c.Param("student_id")

		var studentID int16

		// Convertir id a int16
		fmt.Sscan(idStudent, &studentID)

		var calificaciones[] calificaciones


		if err := db.Table("calificaciones").Select("id, grade, subjectsId").Where("studentId = ?", studentID).Scan(&calificaciones).Error; err != nil {
			c.JSON(404, gin.H{"error": "Materias no encontradas"})
			return
		}
		c.JSON(200, calificaciones)

	})

	router.Run(":8000")
}
