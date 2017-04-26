package main

import "fmt"
import "math/big"

// функции для разложения числа в непрерывную дробь
func convergents(cf [100000]*big.Int) (res_p, res_q [100000]*big.Int) {
	zero := new(big.Int)
	fmt.Sscan("0", zero)
	r := new(big.Int)
	s := new(big.Int)
	p := new(big.Int)
	q := new(big.Int)
	//kek := new(big.Int);
	var i int64
	i = 0
	fmt.Sscan("0", r)
	fmt.Sscan("1", s)
	fmt.Sscan("1", p)
	fmt.Sscan("0", q)
	for _, c := range cf {
		if c == nil {
			break
		}

		//p, q, c*p+r, c*q+s;
		res_p[i] = new(big.Int).Mul(c, p)
		res_q[i] = new(big.Int).Mul(c, q)

		res_p[i] = new(big.Int).Add(res_p[i], r)

		res_q[i] = new(big.Int).Add(res_q[i], s)
		r = new(big.Int).Add(p, zero)
		s = new(big.Int).Add(q, zero)
		p = new(big.Int).Add(res_p[i], zero)
		q = new(big.Int).Add(res_q[i], zero)

		i = i + 1
	}
	return res_q, res_p
}

func contfrac(q1, p1 *big.Int) (n [100000]*big.Int) {
	zero := new(big.Int)
	kek := new(big.Int)
	fmt.Sscan("0", zero)
	q := new(big.Int).Add(q1, new(big.Int).SetInt64(0))
	p := new(big.Int).Add(p1, new(big.Int).SetInt64(0))
	i := 0
	for q.Cmp(zero) > 0 {
		n[i] = new(big.Int).Div(p, q)
		kek = new(big.Int).Add(q, new(big.Int).SetInt64(0))
		q = new(big.Int).Sub(p, new(big.Int).Mul(q, n[i]))
		p = new(big.Int).Add(kek, new(big.Int).SetInt64(0))
		i = i + 1
	}
	return n
}

//Public Shared Function getPAndQ(n As BigInteger, e As BigInteger, fraction As ShortFraction) As ShortFraction
//            Dim k = fraction.p
//            Dim d = fraction.q
//            Dim psi = (e * d - 1) / k
//            Dim a = n + 1 - psi
//            Dim p = (a + BigMath.nthRoot(BigInteger.Pow(a, 2) - 4 * n, New BigInteger(2), New BigInteger())) / 2
//            Dim q = n / p
//            Return New ShortFraction With {.p = p, .q = q}
//        End Function
func getP(n, e, k, d *big.Int) (p *big.Int) {
	psi := new(big.Int).Sub(new(big.Int).Mul(e, d), new(big.Int).SetInt64(1))
	psi = new(big.Int).Div(psi, k)
	a := new(big.Int).Add(n, new(big.Int).SetInt64(1))
	a = new(big.Int).Sub(a, psi)
	kek := new(big.Int).Exp(a, new(big.Int).SetInt64(2), new(big.Int).SetInt64(0))
	kek = new(big.Int).Sub(kek, new(big.Int).Mul(n, new(big.Int).SetInt64(4)))
	kek = new(big.Int).Sqrt(kek)
	kek = new(big.Int).Add(kek, a)

	return new(big.Int).Div(kek, new(big.Int).SetInt64(2))
}

func main() {
	isE := new(big.Int)
	isN := new(big.Int)
	fmt.Sscan("1549128338146508293223128499101054696313591942059932730800245249683705390034596418965793548642658175614765737927271590675064597489123680803300588901534672812795047358187078197056960812301180897331336957518573990892157661234127475641902614398950649137364395227097270289499857256901831284759505242409209429401015047362839366545111661834218537929545154400133239196388247439765240764365392138197482770688668849754650629496820286930014736792050094875604644720080445018044439579507719826400765915124888058455641742003549493191503225508792852888789540340356843881329892655060864644777029412479848018221637677182108001654769", isE)
	fmt.Sscan("23664291872557318092084817496336106189031813559154555271330300903091440071039407435602281123451171726918774041257320301077825689837627095652947164559154016070402883844423719002200739478361354284012340114355277858817821634935185323225536175290442064425643341389634496887571778007923154677319378822648457934490222209358844628456114137916164647071134603664613289537256772630677903109644072947845276836480195383440819617427032363119007597287156781998655267888746187281166013890291098423976156612854872748803650331274897071298637951268753688679157189843636217652149896806895238025221643485869246628474381683593688609814203", isN)
	myM := new(big.Int)
	myM2 := new(big.Int)
	isC := new(big.Int)
	fmt.Sscan("16843009", myM)
	fmt.Println("m = ", myM, " ^^ ", isE, " mod ", isN)
	isC.Exp(myM, isE, isN)
	fmt.Println("m = ", myM, " ^^ ", isE, " mod ", isN)

	//по теореме Винера считаем ограничение для секретной экспоненты D
	limitD := new(big.Int)
	fmt.Sscan("3", limitD)
	limitD = limitD.Sqrt(limitD.Sqrt(limitD.Div(isN, limitD)))
	fmt.Println("Find D < 1/3*N^^1/4: D < ", limitD)

	// раскладываем число E/N в непрерывную дробь и проверяем знаменатель
	// каждой подходящей дроби: не является ли он секретным ключом
	n_arr := contfrac(isE, isN)
	arr_p, arr_q := convergents(n_arr)
	var curr_p *big.Int
	var p, q *big.Int

	isC.Exp(myM, isE, isN)

	fmt.Println("m = ", myM, " ^^ ", isE, " mod ", isN)
	fmt.Println("isC = ", isC)
	for i := range arr_p {
		//		fmt.Println(i)
		curr_p = arr_p[i]
		//		curr_q = arr_q[i]
		if arr_q[i] == nil || arr_q[i].Cmp(limitD) > 0 {
			break
		}
		// С = М^^E mod N
		// M = C^^D mod N
		myM2 = new(big.Int).Exp(isC, arr_q[i], isN)
		if myM.Cmp(myM2) == 0 {
			p = getP(isN, isE, curr_p, arr_q[i])
			q = new(big.Int).Div(isN, p)
			fmt.Println("d = ", arr_q[i])
			fmt.Println("p = ", q)
			fmt.Println("q = ", p)
			break
		}
	}
}
