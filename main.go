package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var contadorProductos int = 0
var idVenta = 0

// ====================================
// Definiciones de estructuras
// ====================================

// Definición del tipo de dato Sucursal
type Sucursal struct {
	Direccion      string
	NIT            string
	NumeroTelefono string
	Empleados      []Empleado
	Ventas         []Venta
}

// Definición del tipo de dato Producto
type Producto struct {
	Nombre   string
	Cantidad int
	Precio   float64
}

// Definición del tipo de dato Empleado
type Empleado struct {
	Nombre       string
	Edad         int
	Puesto       string
	Salario      float64
	Departamento string
}

// Definición del tipo de dato Venta
type Venta struct {
	Id        int
	Productos []Producto
	Precio    float64
	Cliente   string
	Empleado  string
	Fecha     time.Time
}

//====================================

// Funcion para crear la sucursal A
func createSucursalA() Sucursal {
	// Llenar la Sucursal A con empleados
	sucursalA := Sucursal{
		Direccion:      "Sucursal A - Calle Principal",
		NIT:            "123456789",
		NumeroTelefono: "555-1111",
		Empleados:      fillEmployeeA(),
	}
	return sucursalA
}

// Función para llenar una lista de empleados para la Sucursal A
func fillEmployeeA() []Empleado {
	empleados := []Empleado{
		{
			Nombre:       "Cajero 1 (Sucursal A)",
			Edad:         25,
			Puesto:       "Cajero",
			Salario:      25000.0,
			Departamento: "Ventas",
		},
		{
			Nombre:       "Cajero 2 (Sucursal A)",
			Edad:         28,
			Puesto:       "Cajero",
			Salario:      26000.0,
			Departamento: "Ventas",
		},
		{
			Nombre:       "Despachador (Sucursal A)",
			Edad:         30,
			Puesto:       "Despachador",
			Salario:      28000.0,
			Departamento: "Ventas",
		},
		{
			Nombre:       "Administrador (Sucursal A)",
			Edad:         40,
			Puesto:       "Gerente de Tienda",
			Salario:      45000.0,
			Departamento: "Administración",
		},
		{
			Nombre:       "Vigilante (Sucursal A)",
			Edad:         35,
			Puesto:       "Vigilante",
			Salario:      22000.0,
			Departamento: "Seguridad",
		},
	}
	return empleados
}

// Muestra la informacion de la sucursal
func infoSucursal(sucursal Sucursal) {
	// Mostrar información de las sucursales y empleados
	fmt.Println("Información de la Sucursal A:")
	fmt.Println("Dirección:", sucursal.Direccion)
	fmt.Println("NIT:", sucursal.NIT)
	fmt.Println("Número de Teléfono:", sucursal.NumeroTelefono)
	fmt.Println("Empleados:")
	for _, empleado := range sucursal.Empleados {
		fmt.Println("  Nombre:", empleado.Nombre)
		fmt.Println("  Edad:", empleado.Edad)
		fmt.Println("  Puesto:", empleado.Puesto)
		fmt.Println("  Salario:", empleado.Salario)
		fmt.Println("  Departamento:", empleado.Departamento)
		fmt.Println()
	}
}

// Nueva venta
func newSell(producto *map[int]Producto, sucursal *Sucursal) bool {
	var index int
	var id int
	var precioVenta float64 = 0
	var nombreProducto string
	var productosVenta []Producto
	var cantidad int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Cajero: ")
	// Muestra los cajeros de la sucursal
	for i := 0; i < len(sucursal.Empleados); i++ {
		if sucursal.Empleados[i].Puesto == "Cajero" {
			fmt.Println(i+1, " Nombre: ", sucursal.Empleados[i].Nombre)
		}
	}
	fmt.Print("Venta realizada por: (indice) ")
	fmt.Scanln(&index)
	// Validamos el indice del cajero, si es 0 cancelamos la venta
	if index <= 0 || index > len(sucursal.Empleados) {
		return false
	}

	cajero := sucursal.Empleados[index-1].Nombre
	// Agregamos productos a la venta
	for {
		fmt.Println("Ingrese el id de producto, ingrese 0 para cerrar la agregacion de productos")
		fmt.Scanln(&id)
		if id == 0 {
			break
		}
		_, existe := (*producto)[id]
		// Comprobamos que el producto exista
		if !existe {
			fmt.Println("El producto no existe!!")
		} else {
			fmt.Println((*producto)[id])
			nombreProducto = (*producto)[id].Nombre
			fmt.Println("Ingrese la cantidad: ")
			fmt.Scanln(&cantidad)
			// Validamos que exista la cantidad especificada
			if cantidad > (*producto)[id].Cantidad || cantidad <= 0 {
				fmt.Println("Cantidad de productos inexistente en stock")
				continue
			}
			// Guardamos los productos de la venta
			precioProducto := float64(cantidad) * (*producto)[id].Precio
			productosVenta = append(productosVenta, Producto{nombreProducto, cantidad, precioProducto})
			precioVenta += precioProducto
		}

	}
	// Verificamos si se registro un producto
	if len(productosVenta) > 0 {
		var cliente string
		var confirm string

		fmt.Println("Nombre de cliente: ")
		scanner.Scan()
		cliente = scanner.Text()
		// Creamos el tipo de dato venta con los datos ingresados
		venta := Venta{idVenta + 1, productosVenta, precioVenta, cliente, cajero, time.Now()}
		for {
			showInfoSell(venta)

			// Confirmamos la venta
			fmt.Println("Es correcto su registro? (s/n)")
			fmt.Scanln(&confirm)

			if !(confirm == "s") {
				var op string
				// Llamamos a la modificacion de la venta
				fmt.Println("Desea editar la venta? (s/n)")
				fmt.Scanln(&op)
				if op == "s" {
					if !editSell(&venta, *producto) {
						fmt.Println("Has seleccionado la opcion de rehacer la venta")
						return false
					}
				} else {
					fmt.Println("Venta cancelada")
					return false
				}
			} else {
				for i := 0; i < len(venta.Productos); i++ {
					claveProducto := findProduct(venta.Productos[i].Nombre, *producto)
					productoActual := (*producto)[claveProducto]
					productoActual.Cantidad -= venta.Productos[i].Cantidad
					(*producto)[claveProducto] = productoActual
				}
				sucursal.Ventas = append(sucursal.Ventas, venta)
				return true
			}

		}
	} else {
		return false
	}
}

// Mostrar recibo
func showInfoSell(venta Venta) {
	fmt.Println("===========================================================")
	fmt.Printf("Cliente: %s\n", venta.Cliente)
	fmt.Printf("Empleado: %s\n", venta.Empleado)
	fmt.Printf("Precio Total: %.2f\n", venta.Precio)
	fmt.Printf("Fecha: %s\n", venta.Fecha.Format("2006-01-02 15:04:05"))
	fmt.Println("Productos:")
	// Ciclo para mostrar los productos en el array Productos de Venta
	for _, producto := range venta.Productos {
		fmt.Printf("  Nombre del Producto: %s\n", producto.Nombre)
		fmt.Printf("  Cantidad: %d\n", producto.Cantidad)
		fmt.Printf("  Precio Unitario: %.2f | Total: %.2f \n\n", producto.Precio/float64(producto.Cantidad), producto.Precio)
	}
	fmt.Println("===========================================================")
}

func findProduct(nombreProducto string, productos map[int]Producto) int {
	for i := 1; i <= contadorProductos; i++ {
		if nombreProducto == productos[i+10000].Nombre {
			return i + 10000
		}
	}
	return -1
}

// Modificar el registro de la venta (Antes de confirmarla)
func editSell(venta *Venta, producto map[int]Producto) bool {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for {
		// Mostrar la información actual de la venta
		fmt.Println("Información Actual de la Venta:")
		fmt.Printf("Cliente Actual: %s\n", venta.Cliente)
		fmt.Printf("Cajero Actual: %s\n", venta.Empleado)
		fmt.Println("Productos Actuales:")
		for i, producto := range venta.Productos {
			fmt.Printf("%d. Nombre del Producto: %s, Cantidad: %d, Precio Unitario: %.2f, Valor aportado: %.2f\n", i+1, producto.Nombre, producto.Cantidad, producto.Precio/float64(producto.Cantidad), producto.Precio)
		}
		fmt.Println("Precio total: ", venta.Precio)
		fmt.Println("======================================")
		fmt.Println("Opciones:")
		fmt.Println("1. Modificar Cliente")
		fmt.Println("2. Modificar Cajero")
		fmt.Println("3. Agregar Producto")  // agregar vistas dinamicas
		fmt.Println("4. Eliminar Producto") // agregar vistas dinamicas
		fmt.Println("5. Editar Producto")   // solucionar error index out of bounds
		fmt.Println("6. Confirmar Venta")
		fmt.Println("7. Rehacer venta")
		fmt.Print("Seleccione una opción (1-6): ")

		scanner.Scan()
		input = scanner.Text()

		switch input {
		case "1": // Modificar cliente
			fmt.Print("Nuevo Cliente: ")
			scanner.Scan()
			venta.Cliente = scanner.Text()
		case "2": // Modificar cajero
			fmt.Print("Nuevo Cajero: ")
			scanner.Scan()
			venta.Empleado = scanner.Text()
		case "3": // Agregar producto
			var clave int
			var cantidad int
			fmt.Print("Ingrese la clave del nuevo producto: ")
			fmt.Scanln(&clave)
			_, existe := producto[clave]
			if !existe {
				fmt.Println("Producto no válido")
			} else {
				fmt.Println(producto[clave])
				fmt.Print("Ingrese la cantidad del nuevo producto: ")
				fmt.Scanln(&cantidad)

				if cantidad > 0 {
					precioProducto := float64(cantidad) * producto[clave].Precio
					venta.Productos = append(venta.Productos, Producto{producto[clave].Nombre, cantidad, precioProducto})
					venta.Precio += precioProducto
					fmt.Println("Producto agregado")
				} else {
					fmt.Println("Cantidad inválida.")
				}
			}
		case "4": // Eliminar producto
			var clave int
			for i := 0; i < len(venta.Productos); i++ {
				fmt.Println(i+1, venta.Productos[i].Nombre)
			}
			fmt.Print("Ingrese el id de venta del producto a eliminar: ")
			fmt.Scanln(&clave)
			clave = clave - 1
			if clave >= 0 && clave < len(venta.Productos) {
				precioProducto := venta.Productos[clave].Precio
				venta.Precio -= precioProducto
				nuevoArrayProducto := append(venta.Productos[:clave], venta.Productos[clave+1:]...)
				venta.Productos = nuevoArrayProducto
			} else {
				fmt.Print("El producto no existe en el registro de venta\n\n")
			}

		case "5": // Editar producto
			var clave int
			var nuevaCantidad int
			// Se muestran los productos en la venta
			for i := 0; i < len(venta.Productos); i++ {
				fmt.Println(i+1, venta.Productos[i].Nombre, " | Cantidad: ", venta.Productos[i].Cantidad)
			}
			fmt.Print("Ingrese el id de venta del producto a editar: ")
			fmt.Scanln(&clave)
			clave = clave - 1
			if clave >= 0 && clave < len(venta.Productos) {

				venta.Precio -= venta.Productos[clave].Precio
				fmt.Println(venta.Productos[clave].Nombre, "Cantidad: ", venta.Productos[clave].Cantidad, "Precio total: ", venta.Productos[clave].Precio)
				fmt.Print("Ingrese la nueva cantidad: ")
				fmt.Scanln(&nuevaCantidad)
				if nuevaCantidad > 0 {
					precioUnitario := venta.Productos[clave].Precio / float64(venta.Productos[clave].Cantidad)
					venta.Productos[clave].Cantidad = nuevaCantidad
					venta.Productos[clave].Precio = float64(nuevaCantidad) * precioUnitario
					fmt.Println("Producto editado con éxito.")
					venta.Precio += venta.Productos[clave].Precio
					fmt.Println(venta.Productos[clave].Nombre, "Cantidad: ", venta.Productos[clave].Cantidad, "Precio total: ", venta.Productos[clave].Precio)
				} else {
					fmt.Println("Cantidad inválida.")
				}

			} else {
				fmt.Println("Ingrese un id valido")
			}
		case "6":

			return true // Confirmar la venta y salir del bucle
		case "7":
			return false
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Mostrar Ventas
func showSell(sucursal Sucursal) {
	// Mostrar información de ventas
	for _, venta := range sucursal.Ventas {
		fmt.Println("===================================================")
		fmt.Println("		Numero de venta: ", venta.Id)
		fmt.Println("===================================================")

		fmt.Printf("Cliente: %s\n", venta.Cliente)
		fmt.Printf("Empleado: %s\n", venta.Empleado)
		fmt.Printf("Precio Total: %.2f\n", venta.Precio)
		fmt.Printf("Fecha: %s\n", venta.Fecha.Format("2006-01-02 15:04:05"))
		fmt.Println("Productos:")
		for _, producto := range venta.Productos {
			fmt.Printf("--Nombre del Producto: %s\n", producto.Nombre)
			fmt.Printf("  Cantidad: %d\n", producto.Cantidad)
			fmt.Println("  Precio Unitario: ", producto.Precio/float64(producto.Cantidad))
		}
		fmt.Println()
	}
}

// =======================================
// Opciones de productos
// =======================================
// Crear los productos
func createProducts() map[int]Producto {
	// Crea un mapa de informacion
	productos := map[int]Producto{
		10001: {"Leche", 10, 10000.0},
		10002: {"Pan", 20, 4000.0},
		10003: {"Huevos", 30, 1000.0},
		10004: {"Arroz", 15, 7200.0},
		10005: {"Fideos", 25, 4800.0},
		10006: {"Carne de res", 8, 28000.0},
		10007: {"Pollo", 12, 20000.0},
		10008: {"Pescado", 10, 32000.0},
		10009: {"Aceite de cocina", 5, 14000.0},
		10010: {"Sal", 40, 400.0},
		10011: {"Azúcar", 30, 2000.0},
		10012: {"Harina", 18, 4800.0},
		10013: {"Frutas", 50, 3000.0},
		10014: {"Verduras", 60, 2400.0},
		10015: {"Jabón", 15, 6000.0},
		10016: {"Detergente", 10, 8000.0},
		10017: {"Papel higiénico", 20, 7200.0},
		10018: {"Cepillo de dientes", 25, 3000.0},
		10019: {"Pasta de dientes", 20, 4000.0},
		10020: {"Champú", 15, 12000.0},
	}
	contadorProductos = len(productos)
	//fmt.Println(contadorProductos)
	return productos
}

// Mostrar todos los productos
func showProducts(productos map[int]Producto) bool {
	// Acaba si el mapa esta vacio
	if len(productos) <= 0 {
		return false
	}
	// For que muestra todos los productos
	for id, producto := range productos {
		fmt.Println("id: ", id, "| Nombre: ", producto.Nombre, "| Precio: ", producto.Precio, "| Cantidad: ", producto.Cantidad)
	}
	return true
}

// Agregar un producto
func regProduct(productos *map[int]Producto) bool {
	var nombre string
	var precio float64
	var cantidad int

	fmt.Println("Ingrese el nombre del nuevo producto: ")
	fmt.Scanln(&nombre)
	// Validar precio
	for {
		fmt.Println("Ingrese el precio del nuevo producto: ")
		fmt.Scanln(&precio)
		if precio > 0 && precio < 1000000000000 {
			break
		} else {
			fmt.Println("Ingrese un precio valido")
		}
	}
	// Validar cantidad
	for {
		fmt.Println("Ingrese la cantidad del nuevo producto: ")
		fmt.Scanln(&cantidad)
		if cantidad > 0 && cantidad < 100000000 {
			break
		} else {
			fmt.Println("Ingrese una cantidad valida")
		}
	}
	// Crea el producto del tipo Producto y lo retorna
	productoNuevo := Producto{nombre, cantidad, precio}
	contadorProductos++
	(*productos)[contadorProductos+10000] = productoNuevo
	return true
}

// Eliminar un producto
func delProduct(productos *map[int]Producto) bool {
	var id int
	fmt.Println("Ingrese el id del producto a eliminar")
	fmt.Scanln(&id)

	producto, existe := (*productos)[id]

	if !existe {
		return false
	}
	fmt.Println("Producto ", producto)
	delete(*productos, id)
	contadorProductos--
	return true
}

// =======================================
// Funcion principal
func main() {
	var opcion int = 0
	sucursalA := createSucursalA()
	productos := createProducts()
	fmt.Println("Bienvenido!")
	for {

		fmt.Println("=======================================")
		fmt.Println("Ingrese una opcion para continuar")
		fmt.Println("1. Ver informacion legal de la sucursal")
		fmt.Println("2. Registrar venta")
		fmt.Println("3. Registrar producto")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Mostrar productos en stock")
		fmt.Println("6. Mostrar ventas")
		fmt.Println("7. Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opcion)
		fmt.Println("=======================================")

		switch opcion {

		case 1:

			infoSucursal(sucursalA)

		case 2:

			if newSell(&productos, &sucursalA) {
				fmt.Println("Venta realizada")
				idVenta++
			} else {
				fmt.Println("Venta cancelada")
			}

		case 3:

			// Enviamos un puntero a regProduct para modificar el mapa original
			if regProduct(&productos) {
				fmt.Println("El producto se agrego con exito")
			} else {
				fmt.Println("No se pudo agregar el producto")
			}

		case 4:

			// Enviamos un puntero a delProduct para modificar el mapa original
			if delProduct(&productos) {
				fmt.Println("Producto eliminado con exito")
			} else {
				fmt.Println("No se ha podido eliminar el producto")
			}

		case 5:

			// Mostramos los productos, retorna falso si esta vacio y el operador ! lo convierte en verdadero
			if !showProducts(productos) {
				fmt.Println("El stock esta vacio!")
			}

		case 6:

			showSell(sucursalA)

		case 7:
		default:
			fmt.Println("Opcion no valida")
		}
		// Rompe el ciclo si la opcion es 7
		if opcion == 7 {
			fmt.Println("Dia de trabajo terminado con exito!")
			break
		}
	}

}
