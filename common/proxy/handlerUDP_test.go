package proxy

import "testing"

func TestUdpSend(t *testing.T) {
	t.Run("发包测试", func(t *testing.T) {
		err := udpSend("localhost:8080", []byte("hello server, i am the test"))
		if err != nil {
			t.Fatalf("出现错误%v", err)
		}
	})
	t.Run("大量数测试", func(t *testing.T) {
		err := udpSend("localhost:8080", []byte("2535 \n3850 \n9682 \n8212 \n4343 \n9835 \n383 \n3204 \n7854 \n3802 \n8870 \n2813 \n7062 \n8603 \n5552 \n6771 \n9261 \n6617 \n6503 \n5608 \n1561 \n5377 \n2950 \n2113 \n1427 \n8760 \n5020 \n1350 \n6321 \n3861 \n2053 \n9561 \n8772 \n2759 \n7949"+
			" \n6145 \n1950 \n2863 \n1266 \n3364 \n5562 \n8538 \n3583 \n8711 \n6708 \n8727 \n4141 \n8480 \n1579 \n4889 \n1156 \n6181 \n3458 \n8456 \n1637 \n9631 \n5891 \n3578 \n6556 \n7235 \n5213 \n862 \n6837 \n3109 \n4227 \n8232 \n741 \n7513 \n7880 \n2287 \n9669 \n986 \n7771 \n8415 \n5418 \n4171 "+
			"\n5773 \n7650 \n7655 \n4238 \n8609 \n7633 \n5653 \n623 \n4660 \n2237 \n3091 \n6422 \n3811 \n9847 \n4635 \n7342 \n859 \n1083 \n4993 \n5478 \n2488 \n3709 \n6713 \n6855 2756 \n620 \n6296 \n6862 \n384 \n9139 \n2211 \n5366 \n8644 \n4195 \n7350 \n5305 \n5285 \n6933 \n4331 \n3937 \n1572 \n9107 "+
			"\n7101 \n3316 \n3919 \n5233 \n5538 \n7419 \n4999 \n2047 \n7425 \n578 \n7557 \n2041 \n4181 \n5840 \n9022 \n2561 \n5327 \n4772 \n1361 \n2381 \n4918 \n8190 \n6035 \n4912 \n5999 \n7936 \n4530 \n5673 \n7073 \n1869 \n683 \n1601 \n4899 \n4305 \n1766 \n5065 \n3113 \n5616 \n8864 \n4201 \n7013 \n2315 "+
			"\n5245 \n1062 \n5397 \n9300 \n2428 \n1402 \n2342 \n2032 \n2733 \n7728 \n5112 \n6555 \n7242 \n8634 \n3284 \n3544 \n7046 \n9999 \n2120 \n7188 \n4204 \n4824 \n4294 \n7493 \n7183 \n4394 \n9334 \n3355 \n535 \n6566 \n1771 \n8955 \n8061 \n74 \n5425 \n3448 \n2137 \n4187 \n7099 \n1717 \n9317 \n4901 \n4934 "+
			"\n6056 \n7645 \n2991 \n4456 \n2199 \n4694 \n7840 \n993 \n7027 \n5998 \n7825 \n4476 \n3671 \n1540 \n7157 \n4900 \n2558 \n8388 \n9772 \n5792 \n4763 \n6751 \n7682 \n3766 \n7066 \n2139 \n4721 \n8894 \n1867 \n8750 \n3974 \n4677 \n2565 \n4334 \n6590 \n3731 \n8881 \n5135 \n1281 \n129 \n4428 \n4678 \n187 \n9957 "+
			"\n3724 \n8664 \n4447 \n1637 \n8736 \n4157 \n1322 \n4903 \n1126 \n6199 \n7877 \n3184 \n6414 \n4837 \n5017 \n8192 \n952 \n786 \n9808 \n6013 \n5088 \n1907 \n8114 \n4796 \n7510 \n7654 \n5438 \n6779 \n2668 \n2958 \n8130 \n1069 \n8895 \n4597 \n5192 \n270 \n1976 \n4236 \n9824 \n347 \n7261 \n4388 \n6264 \n9676 \n6244 "+
			"\n6418 \n3463 \n420 \n2905 \n3789 \n623 \n8925 \n1219 \n5950 \n276 \n8704 \n1388 \n8811 \n666 \n7394 \n4553 \n7184 \n1607 \n2097 \n9498 \n7327 \n349 \n8545 \n4163 \n3823 \n3194 \n6031 \n8079 \n86 \n5537 \n1398 \n476 \n2996 \n4858 \n6371 \n6343 \n7615 \n1429 \n5895 \n5157 \n7477 \n3736 \n7627 \n8119 \n2170 \n4937 \n777 "+
			"\n3994 \n8465 \n2348 \n152 \n8187 \n4072 \n189 \n4499 \n4301 \n4366 \n8456 \n9737 \n3311 \n3492 \n6977 \n318 \n6722 \n2577 \n563 \n2652 \n574 \n6088 \n5090 \n8498 \n8317 \n1743 \n7295 \n5621 \n7565 \n1394 \n9649 \n5142 \n3259 \n6906 \n6366 \n2735 \n6882 \n1366 \n6945 \n4143 \n3800 \n7634 \n927 \n838 \n7719 \n4521 \n8286 \n3109 "+
			"\n229 \n904 \n9753 \n8908 \n1864 \n8719 \n8999 \n7159 \n8308 \n4382 \n7303 \n990 \n3139 \n314 \n3996 \n7437 \n9790 \n7605 \n8542 \n6866 \n1604 \n8737 \n47 \n409 \n2364 \n8101 \n4704 \n3891 \n1832 \n3514 \n2420 \n5124 \n1553 \n6789 \n8460 \n7815 \n6758 \n7240 \n405 \n798 \n2868 \n5248 \n2658 \n6599 \n7264 \n9245 \n4220 \n4375 \n1767 "+
			"\n5669 \n9364 \n1165 \n4018 \n351 \n6396 \n4628 \n6786 \n9571 \n5889 \n684 \n5297 \n3385 \n1045 \n8195 \n1264 \n5807 \n6347 \n1649 \n8395 \n9221 \n219 \n1952 \n1495 \n4297 \n2204 \n9303 \n3274 \n6365 \n2570 \n508 \n4307 \n5235 \n7406 \n4915 \n9308 \n7589 \n6724 \n5466 \n3374 \n7712 \n5417 \n680 \n2108 \n3393 \n5432 \n1585 \n2267 \n3939 "+
			"\n9685 \n7049 \n4650 \n7323 \n9782 \n1993 \n3280 \n2239 \n6306 \n9605 \n2442 \n9315 \n7873 \n3986 \n895 \n792 \n9762 \n5993 \n4225 \n8791 \n5859 \n5273 \n4117 \n8113 \n1536 \n5642 \n3227 \n3496 \n2597 \n7660 \n2498 \n2744 \n9331 \n2981 \n5304 \n2678 \n4944 \n143 \n6578 \n3960 \n443 \n4722 \n521 \n1638 \n9834 \n4671 \n4817 \n6419 \n1384 \n6795"+
			" \n9866 \n4534 \n6847 \n3377 \n6007 \n2676 \n5151 \n4194 \n1063 \n6694 \n1113 \n4350 \n6968 \n5474 \n7017 \n7846 \n7733 \n5594 \n1551 \n3539 \n1074 \n3917 \n1468 \n1605 \n1865 \n9801 \n3368 \n5622 \n7041 \n1997 \n8072 \n8140 \n2392 \n6039 \n6467 \n8899 \n4998 \n8438 \n6651 \n3646 \n5020 \n2087 \n3591 \n408 \n8974 \n7828 \n4473 \n8591 \n2736 \n2721"+
			" \n9889 \n4651 \n2081 \n7192 \n3900 \n6573 \n7900 \n7569 \n9626 \n5879 \n184 \n7067 \n3222 \n203 \n7026 \n8631 \n2969 \n7389 \n8486 \n4079 \n4590 \n5121 \n5516 \n9988 \n9212 \n1662 \n6684 \n3005 \n9846 \n7911 \n6304 \n1917 \n8698 \n6772 \n1370 \n2588 \n913 \n218 \n2369 \n3556 \n35 \n6430 \n6641 \n8837 \n364 \n7451 \n7489 \n6297 \n2933 \n8570 \n3221 \n5182"+
			" \n4536 \n3146 \n3483 \n6816 \n4057 \n8959 \n3699 \n6247 \n4566 \n4717 \n6820 \n2669 \n3843 \n1444 \n417 \n4502 \n8746 \n2060 \n9668 \n9599 \n812 \n9810 \n4865 \n8425 \n6258 \n5072 \n2480 \n6509 \n9456 \n4847 \n9791 \n3973 \n2402 \n4965 \n4925 \n1506 \n2408 \n8277 \n2067 \n5710 \n370 \n7441 \n8941 \n4832 \n5458 \n511 \n8238 \n30 \n9912 \n7332 \n4161 \n2232 \n421 "+
			"\n6926 \n4152 \n102 \n4445 \n5742 \n9577 \n8613 \n3602 \n4564 \n7483 \n7359 \n6735 \n8932 \n210 \n3598 \n7516 \n3751 \n6358 \n3360 \n1037 \n474 \n1024 \n7755 \n8257 \n9425 \n3530 \n6827 \n6011 \n8410 \n2278 \n728 \n9369 \n6426 \n8884 \n685 \n4110 \n3666 \n9048 \n6930 \n7769 \n8458 \n8839 \n9115 \n5045 \n9973 \n6974 \n3314 \n9616 \n7455 \n9344 \n7494 \n1145 \n4322 \n6372 "+
			"\n8199 \n2385 \n7113 \n1470 \n6576 \n1265 \n7144 \n1014 \n7833 \n8694 \n7817 \n1186 \n1497 \n5523 \n5118 \n4431 \n1941 \n8454 \n8778 \n3031 \n4549 \n3965 \n821 \n2236 \n2540 \n6678 \n9869 \n2640 \n323 \n851 \n6840 \n3508 \n7947 \n4584 \n4631 \n9487 \n7832 \n3370 \n6438 \n6730 \n3009 \n450 \n4625 \n7760 \n5690 \n8943 \n1994 \n4812 \n9779 \n225 \n9414 \n6995 \n4233 \n3099 \n4318 \n2977 "+
			"\n6000 \n9901 \n8204 \n3349 \n5944 \n2936 \n2903 \n8738 \n1473 \n5997 \n6074 \n4806 \n1285 \n5445 \n9516 \n8279 \n2535 \n6083 \n6321 \n3462 \n1906 \n1581 \n5314 \n9585 \n5504 \n1380 \n5329 \n3798 \n8097 \n1775 \n3744 \n1564 \n3185 \n2531 \n8110 \n2235 \n2567 \n3754 \n4991 \n412 \n3634 \n4097 \n8709 \n5841 \n5332 \n2901 \n834 \n5324 \n5529 \n8173 \n9497 \n5095 \n5623 \n6380 \n7227 \n7223 \n3175 "+
			"\n3613 \n3680 \n5479 \n4834 \n4869 \n3283 \n9987 \n3884 \n1387 \n2695 \n9868 \n6416 \n4875 \n9000 \n6741 \n148 \n5736 \n3147 \n1243 \n4785 \n2160 \n3487 \n4804 \n6479 \n7270 \n1978 \n9932 \n9588 \n2585 \n9314 \n7882 \n1558 \n9003 \n3719 \n1187 \n2822 \n4632 \n5685 \n3298 \n9067 \n8274 \n9905 \n9332 \n2072 \n5986 \n5878 \n7274 \n4789 \n3478 \n4769 \n2617 \n2484 \n9409 \n2478 \n2254 \n9730 \n6299 \n7152"+
			" \n8800 \n8599 \n7893 \n4044 \n6144 \n5484 \n1223 \n7904 \n5816 \n5147 \n9981 \n4908 \n897 "+
			"\n4871 \n2256 \n2171 \n2372 \n4802 \n6177 \n6619 \n9575 \n1800 \n2377 \n2182 \n5036 \n5661 \n9963 \n5394 \n3092 \n8208 \n1724 \n4544 \n1486 \n8583 \n7564 \n5933 \n554 \n4365 \n7866 \n1212 \n6318 \n972 \n5431 \n8198 \n4996 \n9569 \n739 \n3439 \n8189 \n4821 \n24 \n902 \n1344 \n355 \n5766 \n4822 \n9528 \n6493 \n1802 \n4176 \n762 \n9758 \n9202 \n1810 \n4993 \n8708 \n2800 \n9"+
			"761 \n8979 \n9206 \n6857 \n2797 \n2661 \n2781 \n3388 \n4948 \n3104 \n6688 \n9789 \n7837 \n403 \n4293 \n2717 \n8783 \n8373 \n3704 \n5927 \n8777 \n8676 \n6173 \n3205 \n7912 \n7567 \n6983 \n2705 \n7850 \n8764 \n6009 \n5803 \n4862 \n7037 \n9228 \n4782 \n6143 \n4914 \n2324 \n8377 \n569 \n9112 \n6403 \n3130 \n638 \n9170 \n363 \n7799 \n4108 \n2872 \n7990 \n630 \n217 \n6234 \n3669 \n5713 \n20"+
			"57 \n1085 \n8164 \n1015 \n2347 \n5784 \n8751 \n2877 \n2610 \n8479 \n9913 \n8419 \n4494 \n893 \n8541 \n5846 \n6874 \n9329 \n9453 \n3416 \n9384 \n5254 \n3653 \n6176 \n9860 \n9015 \n7757 \n606 \n9005 \n8857 \n1530 \n7961 \n6360 \n5085 \n5728 \n3972 \n560 \n3069 \n9341 \n4583 \n6117 \n7052 \n1277 \n1363 \n7898 \n6856 \n7424 \n6462 \n9283 \n6616 \n7234 \n7072 \n3173 \n5038 \n5276 \n9247 \n4349 \n5082 \n445 \n6986 \n3817 \n5506 "))
		if err != nil {
			t.Fatalf("出现错误%v", err)
		}
	})
}
