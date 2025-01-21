# Proyecto: Implementación de Cifrados en Go

Este proyecto implementa una colección de cifrados clásicos utilizando el lenguaje de programación Go. Cada cifrado incluye una breve descripción y su correspondiente implementación para facilitar su comprensión y uso.

## Cifrados Implementados

### 1. Cifrado de Sustitución
Este método reemplaza cada letra del texto original con otra letra según una clave predefinida. Es uno de los cifrados más antiguos y simples.

### 2. Cifrado Caesar
Desplaza cada letra del texto original un número fijo de posiciones en el alfabeto. Por ejemplo, un desplazamiento de 3 convierte "A" en "D".

### 3. Cifrado Atbash
Es un tipo de cifrado de sustitución donde el alfabeto se invierte. La primera letra se convierte en la última, la segunda en la penúltima, y así sucesivamente.

### 4. Cifrados Afines
Utiliza una función lineal para transformar cada letra del texto. La función tiene la forma **E(x) = (a*x + b) mod m**, donde `a` y `b` son constantes y `m` es el tamaño del alfabeto.

### 5. Cifrado ROT 13
Una variante del cifrado Caesar con un desplazamiento fijo de 13 posiciones. Debido al alfabeto de 26 letras, aplicar ROT 13 dos veces devuelve el texto original.

### 6. Cifrado Vigenére
Un cifrado polialfabético que utiliza una palabra clave para determinar el desplazamiento de cada letra. La clave se repite hasta que cubra todo el texto.

### 7. Cifrados de Transposición
Reorganiza las letras del texto original según un patrón definido. Por ejemplo, escribir el texto en una cuadrícula y leerlo por columnas.

### 8. Cifrado de Valla (Rail Fence)
Un cifrado de transposición que organiza el texto en una estructura de "valla" (zigzag) y luego lo lee fila por fila.

### 9. Cifrado de Libro o Ottendorf
El texto cifrado consiste en referencias a palabras o letras específicas en un libro previamente acordado, identificadas por página, línea y palabra.

### 10. Cifrado Playfair
Un cifrado de sustitución por pares que utiliza una tabla 5x5 creada con una clave. Cada par de letras en el texto se sustituye según las posiciones en la tabla.

## Características del Proyecto
- **Lenguaje:** Go (Golang)
- **Estructura Modular:** Cada cifrado está implementado como un módulo independiente.
- **Documentación:** Cada implementación incluye comentarios que explican el funcionamiento del cifrado.
- **Pruebas Unitarias:** Se incluyen pruebas para validar cada cifrado.

## Requisitos
- Go 1.18 o superior

## Instalación
1. Clona este repositorio:
   ```bash
   git clone https://github.com/ManuelR11/Cifrados_Historia
   ```
2. Navega al directorio del proyecto:
   ```bash
   cd cifrados-en-go
   ```
3. Ejecuta el proyecto:
   ```bash
   go run main.go
   ```

## Uso
Cada cifrado puede ser ejecutado de manera independiente mediante comandos en la línea de comandos. Ejemplo:
```bash
go run main.go --cifrado caesar --texto "Hola Mundo" --clave 3
```

## Contribuciones
Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request con tus mejoras o nuevas implementaciones de cifrados.


---

¡Explora los fundamentos de la criptografía clásica con este proyecto en Go!

