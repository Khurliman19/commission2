package main

func ExampleMonthlyPayment() {
	p := Product{
		ID:     1,
		Name:   "Test Phone",
		Price:  2499000,
		Months: 12,
	}

	result := monthlyPayment(p)

	if result != 208250 {
		panic("unexpected result")
	}

}
