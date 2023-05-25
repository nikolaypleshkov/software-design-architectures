# Burger Restaurant
 This code represents a simple burger restaurant implemented in Go. It showcases the usage of different design patterns such as Singleton, Command, Factory, and Decorator.

### Design Patterns
- ***Singleton Pattern (Restaurant)***: The `restaurant` package demonstrates the **Singleton pattern**. The `Restaurant` struct ensures that there is only one instance of the restaurant.

- ***Command Pattern (Cashier and Chef)***: The `cashier` and `chef` packages utilize the **Command pattern**. The `Cashier` acts as a command invoker, taking customer orders and creating `Order` objects. The `Chef` acts as a command receiver, processing the orders and preparing the burgers.

- ***Factory Pattern (Kitchen and Burgers)***: The `kitchen` and `burgers` packages implement the **Factory** pattern. The `kitchen` package acts as a factory by providing the `GetBurger` function to create different types of burgers (King Burger, Classic Burger, and Vegetarian Burger).

- ***Decorator Pattern (Adding Sauces to Burgers)***: The `sauceDecorator` package showcases the **Decorator pattern**. The `AddSauce` function allows adding different types of sauces (Garlic Sauce, Ketchup or Mustard) to the burgers dynamically.

### Running the Code
To run the code:

1. Clone or download the repository to your local machine.

2. Ensure that you have Go installed on your machine.

3. Navigate to the root directory of the project.

4. Execute the following command to run the code:

```bash
go run main.go
```
The program will simulate the burger restaurant workflow, where the cashier takes orders and communicates them to the chef. The chef then prepares the burgers with the selected sauces, and the prepared burgers are served.

You will see the output of the program in the console, displaying the details of the prepared and served burgers.

That's it! You have successfully run the burger restaurant simulation using Go and observed the different design patterns in action.