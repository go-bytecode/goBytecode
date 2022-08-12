package main

type Value interface{}

type Gocodeobj struct {
	co_argcount  int    /* #arguments, except *args */
	co_nlocals   int    /* #local variables */
	co_stacksize int    /* #entries needed for evaluation stack */
	co_code      *Value /* instruction opcodes */
	co_consts    *Value /* list (constants used) */
	co_names     *Value /* list of strings (names used) */
	co_varnames  *Value /* tuple of strings (local variable names) */
}

/*
ox00(000): <0>
ox01(001): POP_TOP
ox02(002): ROT_TWO
ox03(003): ROT_THREE
ox04(004): DUP_TOP
ox05(005): DUP_TOP_TWO
ox06(006): <6>
ox07(007): <7>
ox08(008): <8>
ox09(009): NOP
ox0A(010): UNARY_POSITIVE
ox0B(011): UNARY_NEGATIVE
ox0C(012): UNARY_NOT
ox0D(013): <13>
ox0E(014): <14>
ox0F(015): UNARY_INVERT
ox10(016): <16>
ox11(017): <17>
ox12(018): <18>
ox13(019): BINARY_POWER
ox14(020): BINARY_MULTIPLY
ox15(021): <21>
ox16(022): BINARY_MODULO
ox17(023): BINARY_ADD
ox18(024): BINARY_SUBTRACT
ox19(025): BINARY_SUBSCR
ox1A(026): BINARY_FLOOR_DIVIDE
ox1B(027): BINARY_TRUE_DIVIDE
ox1C(028): INPLACE_FLOOR_DIVIDE
ox1D(029): INPLACE_TRUE_DIVIDE
ox1E(030): <30>
ox1F(031): <31>
ox20(032): <32>
ox21(033): <33>
ox22(034): <34>
ox23(035): <35>
ox24(036): <36>
ox25(037): <37>
ox26(038): <38>
ox27(039): <39>
ox28(040): <40>
ox29(041): <41>
ox2A(042): <42>
ox2B(043): <43>
ox2C(044): <44>
ox2D(045): <45>
ox2E(046): <46>
ox2F(047): <47>
ox30(048): <48>
ox31(049): <49>
ox32(050): <50>
ox33(051): <51>
ox34(052): <52>
ox35(053): <53>
ox36(054): STORE_MAP
ox37(055): INPLACE_ADD
ox38(056): INPLACE_SUBTRACT
ox39(057): INPLACE_MULTIPLY
ox3A(058): <58>
ox3B(059): INPLACE_MODULO
ox3C(060): STORE_SUBSCR
ox3D(061): DELETE_SUBSCR
ox3E(062): BINARY_LSHIFT
ox3F(063): BINARY_RSHIFT
ox40(064): BINARY_AND
ox41(065): BINARY_XOR
ox42(066): BINARY_OR
ox43(067): INPLACE_POWER
ox44(068): GET_ITER
ox45(069): STORE_LOCALS
ox46(070): PRINT_EXPR
ox47(071): LOAD_BUILD_CLASS
ox48(072): YIELD_FROM
ox49(073): <73>
ox4A(074): <74>
ox4B(075): INPLACE_LSHIFT
ox4C(076): INPLACE_RSHIFT
ox4D(077): INPLACE_AND
ox4E(078): INPLACE_XOR
ox4F(079): INPLACE_OR
ox50(080): BREAK_LOOP
ox51(081): WITH_CLEANUP
ox52(082): <82>
ox53(083): RETURN_VALUE
ox54(084): IMPORT_STAR
ox55(085): <85>
ox56(086): YIELD_VALUE
ox57(087): POP_BLOCK
ox58(088): END_FINALLY
ox59(089): POP_EXCEPT
ox5A(090): STORE_NAME
ox5B(091): DELETE_NAME
ox5C(092): UNPACK_SEQUENCE
ox5D(093): FOR_ITER
ox5E(094): UNPACK_EX
ox5F(095): STORE_ATTR
ox60(096): DELETE_ATTR
ox61(097): STORE_GLOBAL
ox62(098): DELETE_GLOBAL
ox63(099): <99>
ox64(100): LOAD_CONST
ox65(101): LOAD_NAME
ox66(102): BUILD_TUPLE
ox67(103): BUILD_LIST
ox68(104): BUILD_SET
ox69(105): BUILD_MAP
ox6A(106): LOAD_ATTR
ox6B(107): COMPARE_OP
ox6C(108): IMPORT_NAME
ox6D(109): IMPORT_FROM
ox6E(110): JUMP_FORWARD
ox6F(111): JUMP_IF_FALSE_OR_POP
ox70(112): JUMP_IF_TRUE_OR_POP
ox71(113): JUMP_ABSOLUTE
ox72(114): POP_JUMP_IF_FALSE
ox73(115): POP_JUMP_IF_TRUE
ox74(116): LOAD_GLOBAL
ox75(117): <117>
ox76(118): <118>
ox77(119): CONTINUE_LOOP
ox78(120): SETUP_LOOP
ox79(121): SETUP_EXCEPT
ox7A(122): SETUP_FINALLY
ox7B(123): <123>
ox7C(124): LOAD_FAST
ox7D(125): STORE_FAST
ox7E(126): DELETE_FAST
ox7F(127): <127>
ox80(128): <128>
ox81(129): <129>
ox82(130): RAISE_VARARGS
ox83(131): CALL_FUNCTION
ox84(132): MAKE_FUNCTION
ox85(133): BUILD_SLICE
ox86(134): MAKE_CLOSURE
ox87(135): LOAD_CLOSURE
ox88(136): LOAD_DEREF
ox89(137): STORE_DEREF
ox8A(138): DELETE_DEREF
ox8B(139): <139>
ox8C(140): CALL_FUNCTION_VAR
ox8D(141): CALL_FUNCTION_KW
ox8E(142): CALL_FUNCTION_VAR_KW
ox8F(143): SETUP_WITH
ox90(144): EXTENDED_ARG
ox91(145): LIST_APPEND
ox92(146): SET_ADD
ox93(147): MAP_ADD
ox94(148): <148>
ox95(149): <149>
ox96(150): <150>
ox97(151): <151>
ox98(152): <152>
ox99(153): <153>
ox9A(154): <154>
ox9B(155): <155>
ox9C(156): <156>
ox9D(157): <157>
ox9E(158): <158>
ox9F(159): <159>
oxA0(160): <160>
oxA1(161): <161>
oxA2(162): <162>
oxA3(163): <163>
oxA4(164): <164>
oxA5(165): <165>
oxA6(166): <166>
oxA7(167): <167>
oxA8(168): <168>
oxA9(169): <169>
oxAA(170): <170>
oxAB(171): <171>
oxAC(172): <172>
oxAD(173): <173>
oxAE(174): <174>
oxAF(175): <175>
oxB0(176): <176>
oxB1(177): <177>
oxB2(178): <178>
oxB3(179): <179>
oxB4(180): <180>
oxB5(181): <181>
oxB6(182): <182>
oxB7(183): <183>
oxB8(184): <184>
oxB9(185): <185>
oxBA(186): <186>
oxBB(187): <187>
oxBC(188): <188>
oxBD(189): <189>
oxBE(190): <190>
oxBF(191): <191>
oxC0(192): <192>
oxC1(193): <193>
oxC2(194): <194>
oxC3(195): <195>
oxC4(196): <196>
oxC5(197): <197>
oxC6(198): <198>
oxC7(199): <199>
oxC8(200): <200>
oxC9(201): <201>
oxCA(202): <202>
oxCB(203): <203>
oxCC(204): <204>
oxCD(205): <205>
oxCE(206): <206>
oxCF(207): <207>
oxD0(208): <208>
oxD1(209): <209>
oxD2(210): <210>
oxD3(211): <211>
oxD4(212): <212>
oxD5(213): <213>
oxD6(214): <214>
oxD7(215): <215>
oxD8(216): <216>
oxD9(217): <217>
oxDA(218): <218>
oxDB(219): <219>
oxDC(220): <220>
oxDD(221): <221>
oxDE(222): <222>
oxDF(223): <223>
oxE0(224): <224>
oxE1(225): <225>
oxE2(226): <226>
oxE3(227): <227>
oxE4(228): <228>
oxE5(229): <229>
oxE6(230): <230>
oxE7(231): <231>
oxE8(232): <232>
oxE9(233): <233>
oxEA(234): <234>
oxEB(235): <235>
oxEC(236): <236>
oxED(237): <237>
oxEE(238): <238>
oxEF(239): <239>
oxF0(240): <240>
oxF1(241): <241>
oxF2(242): <242>
oxF3(243): <243>
oxF4(244): <244>
oxF5(245): <245>
oxF6(246): <246>
oxF7(247): <247>
oxF8(248): <248>
oxF9(249): <249>
oxFA(250): <250>
oxFB(251): <251>
oxFC(252): <252>
oxFD(253): <253>
oxFE(254): <254>
oxFF(255): <255>

*/
