# Finite field (Trường hữu hạn)

## I. Định nghĩa nhóm

- Nhóm là một tập hợp G và phép toán 2 ngôi * (G, *) thỏa mãn các tính chất sau:
    
    + Tính đóng (closure): Với mọi a, b thuộc G thì ta có a * b thuộc G.
    + Tính kết hợp (Associativity): với mọi a, b, c thuộc G ta có:
    ```(a * b)* c = a * (b * c)```
    + Phần tử đơn vị (identity element): tồn tại một phần tử đơn vị e thuộc G thỏa mãn ```e * a = a * e = a``` với mọi a thuộc G. Nếu tồn tại phần tử đơn vị là duy nhất.
    + Phần tử nghịch đảo (inverse element):với mỗi a thuộc G, tồn tại b thuộc G thỏa mãn ```a * b = b * a = e```

## II. Định nghiã trường hữu hạn

### 1. Finite Field (Trường hữu hạn)
- Trường hữu hạn là một ***tập hữu hạn*** các số (tập F) và hai phép toán cộng (+) và nhân (*) thỏa mãn các tính chất sau: 

    + Các phần tử của F tạo thành một ***nhóm*** với phép toán + với phần tử đơn vị là 0
    + Các phần tử của F ngoại trừ 0 tạo thành một ***nhóm*** với phép toán * với phần tử đơn vị là 1
    + Các phần tử của F cùng với hai phép toán + và * thỏa mãn luật phân phối, tức là:
    ``` 
    a * (b + c) = (a * b) + (a * c) với mọi a, b, c thuộc F.
    ```

- VD: Tập hợp F = {-1, 0, 1} với hai phép toán +, * là một trường hữu hạn.

### 2. Finite Set (Tập hữu hạn)
- Nếu kích thước của tập là p thì các phần tử của tập hợp đó là: 0, 1, 2, ..., p-1.
```
Fp = {0, 1, 2, 3, 4, 5, ..., p-1}
```
- Chú ý: Kích thước của tập hợp luôn lớn hơn 1 và là một số nguyên tố .

### 3. Phép toán lấy module

- Để đảm báo tính đóng của trường hữu hạn khi thực hiện các phép toán cộng, trừ, nhân và chia thì chúng ta thực hiện các phép toán module
- Chúng ta định nghĩa phép cộng trong một tập hữu hạn sử dụng phép toán module như sau:
    + Giả sử F5 = {0, 1, 2, 3, 4} thì
    + 1 + 3 = 4 mod 5
    + 3 + 4 = 2 mod 5
    + ...
- Tương tự đối với các phép toán trừ và nhân.

### 5. Định lý fermat nhỏ 
- Nếu p là một số nguyên tố thì:

     ```n ^ (p-1) = 1 mod p```
- Từ định lý fermat nhỏ có thể giúp chúng ta dễ dàng thực hiện phép chia 
- Giả sử ta muốn tính
    
    ```a/b mod p```

    ```Ta có: a/b = a * b^(-1)```

    ```Mà b^(p-1) = 1 => b^(-1) = b^(-1) * b^(p-1) = b^(p-2)```

    ```Vậy a/b = a * b^(p-2)```

## III. Xây dựng Finite Field với Golang

- Tạo struct finiteField

```
type FiniteField struct {
	num   int32
	prime int32
}
```

- Function khởi tạo FiniteField 
```
func Create(num, prime int32) (error, FiniteField) {
	if num < 0 || num >= prime {
		err := "Input wrong"
		e := errors.New(err)
		return e, FiniteField{}
	}
	return nil, FiniteField{
		num:   num,
		prime: prime,
	}
}
```

- Function cộng
```
func Add(f1, f2 FiniteField) (error, FiniteField) {

	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num + f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}
```

- Function trừ 
```
func Sub(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num - f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}
```

- Function nhân 
```
func Mul(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num * f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}
```

- Function lũy thừa 
```
func Pow(f FiniteField, exp int32) FiniteField {
	n := exp % (f.prime - 1)
	num := powMode(f.num, n, f.prime)
	return FiniteField{
		num:   num,
		prime: f.prime,
	}
}
```

- Function chia
```
func Div(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := powMode(f2.num, f1.prime-2, f1.prime) * f1.num % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}
```


