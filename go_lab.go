package main

import "fmt"
import "math/big"

// функции для разложения числа в непрерывную дробь
func convergents(cf [10000]*big.Int)(res_p, res_q [10000]*big.Int){
    zero := new(big.Int);
    fmt.Sscan("0", zero);
    r := new(big.Int);
    s := new(big.Int);
    p := new(big.Int);
    q := new(big.Int);
    //kek := new(big.Int);
    var i int64;
    i = 0;
    fmt.Sscan("0", r);
    fmt.Sscan("1", s);
    fmt.Sscan("1", p);
    fmt.Sscan("0", q);
    for c := range cf {
        fmt.Println(i);
        fmt.Println("c == ", cf[i]);
    	if c == nil {
    	   break;
    	}
        r = new(big.Int).Add(p, zero);
        s = new(big.Int).Add(q, zero);

        fmt.Println("r == ", r);
        //kek = new(big.Int).Mul(c, q);
        //q.Add(kek, s);
        //p, q, c*p+r, c*q+s;
        res_p[i] = new(big.Int);
        res_q[i] = new(big.Int);
        res_p[i].Mul(cf[i], p);
        res_p[i].Add(res_p[i], r);
        fmt.Println("p == ", p, " res_p[i] = ", res_p[i]);

	res_q[i].Mul(cf[i], q);
       res_q[i].Add(res_q[i], s);
        i = i + 1;
    }
    return res_p, res_q;
}

func contfrac(q, p *big.Int) (n [10000] *big.Int) {
    zero := new(big.Int);
    kek := new(big.Int);
    lol := new(big.Int);
    fmt.Sscan("0", zero);
    i := 0;
    for q.Cmp(zero) > 0 {
        fmt.Println("i = ", i);
        fmt.Println(p);
        fmt.Println(q);
        fmt.Println(zero);
        kek.Div(p, q);
        n[i] = kek;
        fmt.Println("n_i = ",n[i]);
	kek.Mul(q, kek);
	lol.Add(q, zero);
        q.Sub(p, kek);
        fmt.Println("q_ ", q);
        p.Add(lol, zero);

        fmt.Println("p = ", p);
        i = i + 1;
    }
    return n
}

func main() {
    isE := new(big.Int);
    isN := new(big.Int);
    fmt.Sscan("1549128338146508293223128499101054696313591942059932730800245249683705390034596418965793548642658175614765737927271590675064597489123680803300588901534672812795047358187078197056960812301180897331336957518573990892157661234127475641902614398950649137364395227097270289499857256901831284759505242409209429401015047362839366545111661834218537929545154400133239196388247439765240764365392138197482770688668849754650629496820286930014736792050094875604644720080445018044439579507719826400765915124888058455641742003549493191503225508792852888789540340356843881329892655060864644777029412479848018221637677182108001654769", isE);
    fmt.Sscan("23664291872557318092084817496336106189031813559154555271330300903091440071039407435602281123451171726918774041257320301077825689837627095652947164559154016070402883844423719002200739478361354284012340114355277858817821634935185323225536175290442064425643341389634496887571778007923154677319378822648457934490222209358844628456114137916164647071134603664613289537256772630677903109644072947845276836480195383440819617427032363119007597287156781998655267888746187281166013890291098423976156612854872748803650331274897071298637951268753688679157189843636217652149896806895238025221643485869246628474381683593688609814203", isN);

    //по теореме Винера считаем ограничение для секретной экспоненты D
    limitD := new(big.Int);
    fmt.Sscan("3", limitD);
    limitD = limitD.Sqrt(limitD.Sqrt(limitD.Div(isN, limitD)));
    fmt.Println("Find D < 1/3*N^^1/4: D < ", limitD);

    myM := new(big.Int);
    myM2 := new(big.Int);
    isC := new(big.Int);
    fmt.Sscan("16843009", myM);

    // раскладываем число E/N в непрерывную дробь и проверяем знаменатель
    // каждой подходящей дроби: не является ли он секретным ключом
    n_arr := contfrac(isE, isN)
    arr_p, arr_q := convergents(n_arr)
    var curr_p *big.Int;
    for i := range arr_p {
        fmt.Println(i);
        curr_p = arr_p[i];
        //curr_q = arr_q[i];
        if arr_q[i] == nil || arr_q[i].Cmp(limitD) > 0 {
            break
        }
        // С = М^^E mod N
         isC.Exp(myM, isE, isN);
        // M = C^^D mod N
         myM2.Exp(isC, arr_q[i], isN);
        if (myM == myM2) {
            fmt.Println("is fraction: ", curr_p,"/", arr_q[i]," Need check denominator: ", arr_q[i],". VALID VALUE")
        } else {
            fmt.Println("is fraction: ", curr_p,"/", arr_q[i]," Need check denominator: ", arr_q[i],". Invalid")
        }
    }
}
