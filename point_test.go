package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

var (
	GermanyNodes = []int{7008268, 5954074, 9120618, 5747785, 6367494, 4381896, 165343, 4695953, 301566, 6781958, 6194766, 2358012, 4692093, 497151, 4355495, 700415, 4390319, 8553348, 6542963, 6634588, 6517688, 5392367, 6190253, 4198881, 504621, 9228869, 7186081, 895834, 5606641, 2647250, 6339514, 691811, 2308165, 10001569, 9120051, 759941, 10073304, 6311612, 6343270, 8179683, 3366117, 6061132, 5643689, 863577, 4931326, 6649953, 294592, 4419734, 4802743, 9567772, 5985694, 468718, 7761320, 5053336, 7260626, 4031757, 4574787, 3596809, 9684347, 164709, 195923, 9834961, 8155072, 7889391, 5200362, 8055855, 8304451, 3560874, 6367318, 5940849, 4633454, 1769665, 7338960, 7345082, 6542540, 1699666, 4264625, 7749655, 10239981, 9265680, 9749596, 6070658, 7983593, 1573825, 5277830, 6966871, 4712734, 7511618, 5665254, 8081312, 2438484, 6431216, 6041290, 2647929, 10208160, 143963, 2291010, 4056536, 5649426, 9047055, 6949841, 4713404, 7111360, 1606682, 735773, 2156000, 1852553, 4645613, 10440701, 3375948, 2775921, 4835420, 1425989, 3591738, 6262150, 7189083, 2183250, 1455662, 9080962, 4939489, 5849163, 2481864, 5358757, 3992900, 8099433, 6030163, 9543396, 3100653, 46963, 6764954, 6741542, 1499985, 1518525, 5777389, 7179979, 6021581, 4558442, 4698864, 10384852, 9110113, 404713, 4015965, 839030, 3539652, 8428891, 171848, 3299486, 6890628, 6617824, 7543977, 7611477, 3287977, 6353455, 10167787, 4457068, 904612, 1125329, 6047234, 2655479, 5740142, 8520248, 334569, 1932239, 7974489, 6232975, 8266989, 5910401, 7383070, 7639188, 906419, 1633652, 4206622, 9425661, 643972, 3654958, 9809412, 9690697, 2206079, 8277645, 5117033, 3990216, 8974041, 5227650, 2746089, 3557062, 127803, 3631098, 5863287, 4719210, 4256235, 8015655, 159987, 7102034, 6788209, 1203170, 6874024, 4575612, 8767789, 2668569, 2569985, 3715654, 393198, 9950700, 9756778, 4282083, 3597652, 5530427, 10165444, 2609535, 3923065, 3748038, 238269, 8862259, 1212190, 5192509, 10063862, 5217308, 6709583, 1871292, 7000930, 7729292, 8338947, 2358378, 10124331, 145863, 10526984, 3662053, 3349217, 7926182, 1962950, 4896276, 8337446, 83303, 4710409, 8901360, 827293, 2015800, 8127920, 2283773, 4786904, 5647141, 5537861, 4039566, 6844087, 538542, 5215683, 10380146, 3521540, 3320306, 3600260, 4355790, 10133601, 3836895, 5575268, 5077219, 2835701, 1422817, 4077361, 6624582, 4221540, 2904403, 4568293, 5892519, 184689, 2476077, 6056339, 109807, 1612422, 3878049, 191945, 10115368, 3764898, 735073, 1606269, 2589551, 2055257, 2949524, 8837422, 4773969, 1356119, 5844509, 8743064, 3449998, 6363927, 1963464, 1983618, 2834810, 9946551, 7951028, 9582813, 1855524, 6821228, 7378037, 3085397, 9599672, 5504082, 3643548, 10565967, 10326036, 7670349, 7211086, 3371299, 8666235, 3342981, 1970318, 8036747, 8855710, 6715344, 9585075, 2995129, 2846975, 5385884, 8406262, 1168258, 5136163, 9398874, 8553238, 9561641, 2487243, 8547976, 6770296, 10585918, 5460042, 4070668, 7268577, 7070856, 2183469, 8121785, 4978500, 9495888, 875326, 810751, 8061641, 5925757, 5352679, 7937646, 3229500, 4451255, 3863865, 4262653, 8164172, 5009125, 1061420, 1066937, 7804651, 4060818, 5304379, 5190842, 9383175, 666907, 8281379, 7159412, 7146625, 2493675, 2189459, 4153487, 1364599, 3256653, 3861415, 5277930, 1579058, 5138291, 4738642, 4338550, 10017557, 10172942, 5972789, 9085324, 6949823, 7671036, 8927855, 159012, 7640734, 982140, 8675521, 7555198, 4797299, 958078, 6439657, 4461366, 7253947, 1865589, 2529756, 6965863, 6353404, 6154048, 3822233, 7550840, 9244707, 8180305, 9881056, 9901476, 6512218, 8187373, 7028468, 2212809, 9609257, 4383017, 7544524, 8469002, 4546514, 6588534, 10041063, 4410506, 2033493, 3978622, 5041533, 8845909, 6292290, 4506460, 3279146, 127584, 9268812, 4898722, 697244, 5342651, 7437235, 1741089, 6156143, 8740165, 6933569, 830385, 3764164, 5559003, 10462610, 1766341, 3784634, 7014785, 4717196, 4466594, 8057635, 7076472, 930882, 6673984, 3461219, 8542327, 1870640, 3191606, 1265695, 10301216, 9822207, 8158439, 9145925, 622894, 2022618, 4486588, 7872226, 3930117, 3135941, 143016, 207481, 5983237, 1844578, 554724, 7186015, 10277961, 9440616, 3497280, 5298482, 9436673, 4985960, 8937542, 7895146, 9350704, 5856783, 10141258, 7916923, 9759776, 1645301, 2663920, 4904605, 6221967, 5800732, 328798, 1343009, 7878278, 5994601, 1380646, 3802789, 3976593, 9754952, 7503363, 7447120, 8949633, 175114, 5369946, 3949817, 7763354, 2556992, 9232443, 9846625, 1374208, 2738282, 2232863, 6635918, 7082431, 188590, 5062859, 9347117, 9522194, 7384317, 3151723, 5790520, 350632, 4746176, 8301449, 10113215, 5739989, 8544128, 10130623, 8272839, 5273325, 6800692, 5221380, 207425, 2870572, 9384991, 4228119, 3089964, 10211044, 1812681, 9673233, 4317553, 10417209, 6968955, 190930, 5595999, 10537677, 4432445, 6675223, 4348551, 2486124, 7560046, 7875737, 2224522, 5334403, 2900059, 5173704, 4829428, 354958, 4612257, 230953, 3438442, 4916933, 433135, 5134179, 9567579, 3714532, 160789, 10223779, 1574062, 10126314, 406464, 2925625, 63545, 6713089, 2840999, 1770619, 2499861, 5841858, 3655558, 2065718, 2699941, 7833699, 4085746, 9797508, 10258809, 4007056, 3280313, 908962, 2362206, 3374847, 2692578, 4791371, 10339741, 4447057, 27326, 9228746, 7317886, 3813225, 1355246, 6926538, 344915, 7854657, 8070652, 6348184, 8793219, 3959817, 2959288, 5461193, 7251089, 2051443, 8028291, 5587163, 8107348, 7387826, 3904565, 9552186, 9012617, 9644507, 6244858, 10483042, 4947804, 9107840, 2835187, 666236, 3902321, 3755086, 8854595, 6603169, 2228170, 3278122, 4322704, 2783388, 2927936, 2077179, 3664080, 2756663, 8615262, 2650330, 3612977, 362893, 7138806, 4160611, 6130206, 2522191, 2303753, 6150581, 5495525, 5976751, 7914229, 7554878, 9716864, 2027329, 4858963, 705051, 130939, 10227717, 7200193, 7899914, 2108106, 5126168, 10275360, 3480845, 2468193, 3528831, 1637373, 9353445, 2344388, 3094492, 9387507, 1971870, 1450956, 5293219, 3324249, 10072609, 2974883, 5131210, 303382, 8755026, 6597588, 7176437, 9661586, 8823476, 4539105, 6996235, 10323604, 7540604, 10323604, 26878, 10339000, 6754440, 8096435, 5878620, 4636432, 5166363, 8387190, 3389593, 10100787, 9972076, 5161371, 7095930, 2314297, 1532427, 1399265, 8525105, 8094642, 1646511, 2565969, 7569106, 798570, 6149906, 7797629, 8529769, 3925163, 2665077, 4706922, 8452894, 7895970, 9065633, 8929508, 953061, 7489944, 1043017, 1150774, 1791906, 7067140, 8910860, 1615081, 3378362, 5055086, 2895189, 1143132, 1098338, 9091990, 3764440, 1329677, 1095240, 2134596, 376440, 7755346, 4664463, 3968851, 7407147, 1918233, 10486553, 3594027, 3569784, 10355838, 5208748, 1730150, 4719139, 4519465, 7883781, 3966568, 10399161, 1869415, 303580, 663576, 4664071, 1049698, 4973826, 1897608, 1303071, 9744364, 7615826, 7161152, 7639923, 5313002, 3973761, 7531881, 4710220, 7025871, 10332865, 6467444, 9409898, 3299787, 1441993, 4130481, 774919, 3436968, 2243358, 5385280, 1930284, 1265408, 1510728, 5354306, 2513946, 5590283, 3883447, 4917877, 119460, 8988850, 3605413, 1100137, 505859, 9769018, 2898271, 1680550, 2877056, 616522, 7312516, 3219135, 7002941, 8514322, 4623749, 8059209, 5478431, 1948994, 8899630, 856875, 3760740, 548128, 7570585, 4420376, 4680956, 9130435, 1743238, 9951377, 5680356, 3621931, 4808892, 7739623, 5612649, 4652520, 7637168, 8486560, 1169084, 5750253, 8106464, 3322550, 6415269, 4268984, 1577996, 2711791, 2983814, 6507780, 510646, 4858426, 8359975, 7178174, 5905330, 7497618, 8453129, 611675, 2032475, 9266536, 82536, 6186781, 8649640, 3539719, 2137295, 1822976, 383231, 4503728, 7984482, 1589693, 2788632, 6649261, 8262838, 2852931, 6717509, 632232, 6201618, 3083041, 1606088, 730490, 4404797, 9578295, 596238, 3370193, 562873, 9439953, 1690991, 1929476, 9066090, 7611489, 10270426, 3647856, 8568337, 4746105, 6903963, 1624104, 685158, 1885113, 669531, 5564808, 1532774, 1810808, 4996280, 3267069, 2393986, 2049538, 1269357, 3249197, 3255665, 6454868, 5044307, 168909, 3520924, 8461625, 2513134, 4698731, 5312672, 9081962, 7995378, 5420177, 9593553, 7001111, 3781492, 1932856, 9489588, 4857500, 8401967, 8413362, 8644553, 2592055, 4047948, 4832632, 4945914, 1911366, 8465577, 4764362, 3384502, 7935311, 2196238, 5721812, 1244373, 8767292, 5506851, 1593418, 9017268, 2278166, 6003232, 787708, 10060991, 9152340, 5507662, 8370922, 2896772, 7564784, 5083027, 8622140, 1866342, 3685589, 10054899, 9045403, 1457010, 852021, 5754600, 4509911, 6848338, 9175650, 222861, 7881978, 527718, 10420826, 4042795, 5381901, 3519703, 1268879, 4817071, 7434667, 9645389, 832995, 9804398, 2599871, 2127462, 8609151, 8368858, 7572734, 876412, 7370391, 2133627, 4040477, 9763382, 523261, 5377040, 5614690, 10103223, 9466244, 6831062, 1427262, 1981776, 8210088, 4807620, 9553483, 7679431, 2438537, 3045851, 1853221, 1933944, 7682890, 3973954, 6086092, 793886, 5826133, 2117937, 7684296, 656902, 1172996, 7658411, 3711352, 6320604, 10256151, 8855072, 3984934, 7842045, 9958278, 9896249, 6762062, 8228823, 6135488, 4555216, 331512, 4437829, 956058, 4659171, 1214911, 2552205, 327256, 7185600, 8584063, 7355530, 5412337, 8273038, 3271855, 1891003, 8611193, 6249420, 1565554, 8335286, 4748241, 1671337, 1930733, 4973167, 4744365, 2327872, 9111959, 2235982, 7148776, 6260433, 9278565, 6261876, 4031952, 285483, 156148, 285940, 10341774, 9652905, 9940143, 1755487, 10504747, 5057875, 6727666, 1965436, 44485, 2878194, 3954976, 1078210, 6845094, 8505807, 2196879, 371519, 1612941, 1097831, 9253703, 3142851, 3374040, 8286778, 3364038, 3076438, 6408258, 8718115, 3972224, 4208595, 575319, 9199426, 9165044, 5070543, 6031638, 8187112, 9585203, 4381902, 8312925, 2546974, 3916399, 986561, 3978830, 9771269, 8857763, 218291, 8766291, 10072214, 10122280, 2233241, 2034473, 7985289, 3560906, 8044769, 746533, 6783033, 6064616, 1183112, 8136889, 2346671, 7990662, 9758621, 7400485, 4455977, 8153931, 9004165, 7715283, 4880073, 8761831, 4699029, 1920402, 2986723, 6251237, 9739403, 1063315, 4696215, 10010994, 1024372, 367787, 7740286, 7264142, 3437219, 3662531, 8368665, 3993060, 7581161, 8088704, 8594492, 5743649, 8695405, 4300209, 1267578, 2498481, 1660065, 4829280, 1655813, 1135893, 3870160, 9351410, 6293744, 585067, 6293744, 9131059, 2667044, 670656, 9374615, 1439894, 9201594, 1962838, 8582745, 679764, 5969809, 7236000, 2259030, 4543476, 8721887, 5680739, 7562234, 2258845, 3435084, 1642192, 6958272, 200966, 6409511, 1356133, 1049881, 6245136, 2536250, 9391957, 9614730, 7732395, 9103538, 7682767, 5429236, 9594997, 4381954, 3138369, 3802409, 3575951, 9684263, 607276, 2847452, 345812, 8575377, 7524529, 525623, 81040, 10036290, 10443308, 8328225, 9856756, 4438495, 4407472, 4828847, 7606612, 4811450, 9850299, 9860552, 8128281, 4931698, 9370095, 3254795, 7323820, 7918587, 6156764, 7363422, 8321952, 9507247, 1126781, 2522399, 6280777, 6026594, 10389500, 2456001, 4707171, 10427528, 5315985, 6008941, 1367604, 10004672, 10384570, 9605561, 28911, 245138, 3492948, 2150274, 9146642, 10068079, 2331237, 2275219, 687241, 1276511, 15691, 4298284, 6091947, 1661852, 10066929, 10379832, 4695709, 462825, 4672484, 5221176, 2004089, 2618557, 10437051, 2995744, 1608811, 8081056, 2514136, 6819999, 8520678, 6635119, 6501595, 1258400, 10286486, 7506545, 3529266, 1052947, 1573824, 3817593, 2043104, 5000913, 9297201, 3545624, 1193398, 4701500, 7443315, 7834661, 270860, 3798558, 6257736, 1757379, 1156173, 2456263, 648469, 4517139, 9699320, 480440, 7047276, 6218549, 6219216, 3795121, 6136909, 8193910, 786310, 3673047, 9250264, 1987243, 9102977, 8719427, 7246610, 10075314, 1090815, 3190141, 8467485, 10161711, 9097527, 4807300, 7324449, 9017895, 2835403, 10390577, 7680897, 1030521, 5218654, 4886655, 4682741, 7652240, 3216064, 3474074, 6827778, 6674460, 1896601, 2696671, 2853286, 3553458, 6378606, 2207584, 9784726, 1496403, 3433135, 10425941, 1248620, 9486325, 2527414, 3101720, 5596533, 2471901, 511387, 8191193, 930885, 6339498, 1182559, 10313300, 509389, 6474815, 9663727, 8130135, 2100548, 255523, 9015771, 9676614, 4494262, 6700578, 7007003, 10110790, 6276777, 9930278, 9704278, 2248472, 6738414, 1481141, 6548422, 3181326, 10117754, 9105016, 7227811, 4483310, 10587768, 7535195, 3118372, 2446099, 10164620, 2296726, 36235, 8987218, 3450508, 3939689, 431325, 3541760, 4257152, 8404060, 774154, 9716006, 3659589, 8069519, 8192553, 5227527, 9098280, 8221694, 9307023, 10014740, 8112489, 2063955, 1888855, 5817439, 10404524, 4955613, 1127339, 9452352, 3366025, 9485194, 7337604, 4071699, 10376791, 3936842, 9891710, 7709467, 9214425, 5846364, 9910005, 9047276, 5050795, 5848953, 9788790, 7976675, 9615584, 6209886, 9123527, 8083870, 9836520, 4430966, 6264239, 959945, 10428317, 3785642, 359542, 8561392, 5056035, 4114284, 3985802, 623361, 4708804, 2688343, 4134935, 6526484, 5817360, 1449524, 10195105, 4604204, 151515, 2388159, 8126684, 6303854, 7080432, 5585913, 8197720, 8827297, 556101, 7661615, 10184282, 3265738, 4212294, 5937254, 9143505, 7482706, 6813600, 6301597, 9667900, 8308578, 7407441, 552227, 8381954, 5649934, 7411991, 6680871, 4788293, 5232881, 7954010, 4491530, 8752463, 3753846, 7253788, 6404036, 89228, 3889321, 5828893, 4351061, 5143457, 8828820, 710185, 2514731, 6481649, 7633186, 9419066, 3422506, 9771760, 9906133, 8062348, 371292, 7411515, 2253002, 4918547, 8494525, 2806697, 6583472, 8195837, 2459566, 2426085, 613499, 8641681, 6427651, 7621003, 3519821, 3272792, 7355954, 3738922, 9788538, 8969624, 5657101, 7670886, 3280880, 3238784, 4655852, 1441315, 9752233, 9037871, 5906910, 3802298, 4721995, 2722443, 9747407, 6128661, 6535160, 1317412, 7198112, 1106953, 500874, 3252603, 7549676, 2383143, 1472498, 1735005, 3910604, 7112822, 8224692, 6081827, 137157, 6636596, 5732292, 7772907, 1032612, 1484138, 1722819, 3588742, 8906616, 10470696, 5920156, 8701452, 301781, 7298813, 10107053, 10024695, 780655, 1892378, 2795666, 5203638, 6939987, 2039138, 3367282, 7256236, 8742011, 2356841, 5158397, 8453466, 3833156, 40674, 8547396, 2940124, 8010214, 10222221, 7308926, 3252736, 1238766, 6842169, 3841631, 8310012, 1054181, 8708602, 2822895, 443774, 4692687, 6507713, 9094494, 9920155, 1815376, 2379690, 4264649, 2856019, 4312215, 7855557, 197398, 4860518, 4955505, 4422490, 7785232, 6059734, 1861679, 5606535, 501608, 4971939, 4982732, 3214285, 4480539, 7113190, 2886983, 7864755, 3243702, 6745997, 3238260, 9392914, 2795072, 5729210, 4159610, 2566706, 4214444, 5170758, 9199797, 8677131, 10137614, 3190880, 6703364, 3393062, 402696, 1181219, 1396762, 9192100, 10070278, 5629390, 2206870, 3508242, 8202128, 9641069, 1978052, 3672314, 5081392, 3041465, 3486392, 352981, 5518127, 2574871, 6264431, 10416132, 5515542, 9500, 6231927, 8023188, 6763423, 6335634, 6769370, 7764710, 5807388, 145797, 7452600, 5103318, 1098592, 5214118, 7411628, 7072339, 4567229, 8563893, 4170122, 5189527, 6603657, 3231455, 3275875, 9227801, 8171981, 2306372, 9326489, 6160397, 2714488, 1556677, 483376, 10495205, 3476663, 7171188, 3654473, 4046427, 3038328, 3487585, 9783933, 1454601, 1488230, 7322436, 8838742, 3006201, 9070897, 8308520, 5207592, 862768, 2023904, 10571192, 1920334, 7763915, 10315447, 164547, 9805365, 3698228, 6232120, 6179996, 5859643, 9361569, 2880776, 10103699, 1832365, 3252724, 279022, 8634504, 9501720, 3691391, 1516472, 4999169, 6896014, 8474965, 5531985, 6460368, 7104929, 8064454, 2203726, 7109968, 3660215, 8005825, 6977382, 4626914, 6849598, 6606589, 602321, 9057689, 4675160, 9680374, 5173225, 7491509, 10082392, 6532432, 9232758, 4729900, 346272, 5623016, 6375059, 7790539, 2719737, 9373457, 9083640, 5625873, 7441523, 2131547, 7311494, 988607, 7102208, 4091168, 8849729, 6357197, 5045470, 128826, 496075, 2623244, 2874056, 10295786, 7945659, 2608629, 2604944, 3487715, 6926810, 2105637, 7711844, 5935101, 9765124, 76963, 3257056, 6204583, 1703834, 3336717, 8067336, 7933503, 5620892, 9157819, 7343326, 9795627, 2832332, 6185765, 4355105, 4646882, 1011757, 3307332, 6865697, 2822010, 9541659, 7593178, 969581, 9540583, 7686475, 2986125, 8103571, 3641011, 8436250, 3207855, 5888078, 3925107, 4690178, 1611150, 4674280, 4809506, 6012528, 8155688, 6118863, 6993540, 10488205, 5650261, 7135308, 7016513, 3075592, 7229620, 8357128, 3817451, 5772210, 7215365, 1831492, 7236755, 7011050, 951634, 7368254, 920846, 3072187, 9971087, 2445356, 3018223, 8989975, 6357085, 4017561, 4626472, 5595672, 6831485, 345267, 2568318, 6116007, 9787561, 6038040, 4223182, 4122631, 7570602, 4960081, 1609190, 1118401, 9459430, 8999422, 1367600, 5035970, 5365072, 2000544, 7388570, 495383, 3647482, 5564235, 9741009, 3967773, 2877607, 2636853, 1882930, 1484422, 7998735, 2678484, 876868, 954739, 9996053, 3183909, 6048726, 4794120, 1005318, 3537499, 6607360, 3140164, 5214961, 4623705, 10420541, 2351281, 4091957, 2531717, 7302689, 9277643, 767539, 7506835, 5718629, 10237591, 5107077, 1800181, 8318930, 7218472, 4386920, 6338979, 10156012, 134340, 3694019, 1189793, 1081460, 3097962, 6141697, 2178548, 1254067, 3062463, 3219876, 9833760, 6033410, 3246654, 2912588, 2841892, 6369115, 406456, 5945640, 2593565, 5474274, 6969454, 2585051, 6508331, 4499938, 7332094, 3829036, 6132745, 1176412, 1566383, 5040268, 2933822, 2164117, 816176, 1309068, 9013137, 7729691, 9291053, 6241130, 665967, 4007993, 9786066, 5771307, 9517370, 10215005, 10452224, 9454039, 2707917, 2022433, 7085775, 8438768, 3541656, 9637460, 5032037, 4216244, 170867, 6548129, 8852925, 1060966, 8331673, 1983954, 5688707, 3568754, 9472112, 5706682, 6161751, 4309909, 7166798, 832393, 7331062, 4130571, 10104395, 5419492, 7512456, 6536413, 1737461, 1772430, 865398, 8704868, 6897377, 8184085, 8000768, 1323760, 1738251, 4317369, 9859564, 4449924, 5881428, 9404658, 1635591, 2560914, 2948519, 731615, 1168343, 3891545, 5765962, 475517, 7063831, 2555960, 5776492, 6383788, 3692414, 335245, 10413070, 4963925, 10546066, 2538949, 7223677, 327871, 10433323, 8915320, 8111160, 4771485, 3404534, 2400644, 2704918, 6545091, 6933101, 5795766, 3956631, 4796894, 2792737, 2596263, 3280163, 1068100, 9678147, 4988261, 10395053, 7800892, 7381999, 9657986, 3251637, 4251377, 19843, 7468518, 10322723, 6650191, 2981934, 1468118, 6485878, 1925123, 1188687, 585742, 5480579, 5587089, 4547125, 10360821, 9197699, 8551519, 7512829, 3245264, 7472969, 7697169, 6005424, 4530664, 4463879, 7615689, 3784069, 207590, 1298677, 9030554, 5603992, 8802010, 2053954, 2918479, 6000342, 7893089, 6516507, 7733744, 10227201, 6222950, 6008637, 3473670, 7943770, 10239563, 5590605, 1927866, 4436758, 5031057, 4681581, 6247227, 10350288, 2449453, 5119814, 8067224, 5079278, 4185826, 4693679, 1600909, 2489739, 5772496, 3658925, 2928638, 3988376, 5294426, 8825049, 3389125, 6162610, 6469292, 609154, 249682, 1116922, 8732903, 2496011, 2869739, 355322, 7306589, 9498646, 9302321, 2617602, 572326, 7548426, 2486129, 3730140, 8564061, 1631460, 2356033, 6875585, 9705924, 8860015, 494840, 6387758, 9048012, 4476248, 1610477, 8341053, 7663243, 1370539, 6871812, 2676882, 10120375, 5487652, 4502649, 8191526, 8181903, 9207762, 5474243, 2262652, 65364, 8999609, 4413726, 9773100, 10332395, 2404098, 3499499, 7995249, 1099743, 10407222, 4201424, 7002816, 10119096, 7671144, 4763726, 7993282, 4808967, 4241919, 4589234, 9992933, 7520011, 7245640, 913163, 9062837, 6868702, 9328339, 6345393, 9888968, 7451213, 6983657, 8649621, 4794419, 4391081, 6045269, 53515, 9689509, 7503545, 5384628, 5952941, 2248792, 10009980, 9266357, 7565853, 6444933, 4895957, 10095721, 1949948, 3123770, 8648630, 9938200, 1223254, 2922950, 7913311, 10011872, 2368425, 8398937, 7755306, 3655, 905702, 6962537, 10176330, 5603461, 3153544, 6540558, 9673664, 10159164, 4175527, 8414365, 9231004, 1180958, 2365453, 8758087, 2144578, 7068819, 6693125, 5674609, 6593129, 8012209, 1594193, 3962537, 3171421, 9034428, 319209, 9110694, 3761624, 8465744, 1357150, 3668343, 4664309, 9260408, 851926, 3902758, 3509591, 3926479, 6214112, 401576, 1924606, 10104170, 8412640, 115712, 4182426, 803448, 3208053, 4994144, 9823494, 444307, 2384939, 8918817, 7583192, 350309, 3954983, 1710796, 3129828, 6615049, 5767925, 7042341, 7186250, 6046231, 2563849, 4622054, 7432422, 10581856, 2166578, 335390, 9720926, 1170777, 755753, 8019978, 460108, 10231195, 1526931, 348113, 8141158, 5700028, 9001520, 4855660, 6278907, 7004192, 4455784, 9908228, 2300412, 8295829, 8962879, 2124230, 6267978, 1934984, 5445605, 8502998, 9212983, 1862997, 394354, 9260326, 6694686, 8216651, 5402307, 5124692, 1341284, 3376627, 9371330, 9321225, 3435486, 525582, 5666807, 9481225, 10526149, 7938555, 6627262, 8854171, 6693251, 3453863, 2689536, 7720926, 3543416, 60712, 8272917, 6259216, 1679557, 6009846, 3886949, 1416645, 5361726, 39983, 8401370, 7474493, 191071, 3646960, 5757158, 3358905, 6790980, 8483048, 309324, 4656913, 6875845, 2916818, 8288908, 2704846, 695216, 4386330, 1380190, 7856105, 10145014, 7158213, 7278933, 5430267, 7279364, 9829677, 9179949, 3106346, 5920696, 4713610, 9071778, 4078247, 4297799, 7662576, 7313179, 2440186, 4727864, 7508451, 3721109, 3329611, 5493005, 7513223, 2631319, 1037944, 10301689, 3942332, 2247867, 10548012, 6804976, 4492610, 2365668, 8902990, 721274, 12614, 5593508, 2132778, 5150612, 9604471, 5581845, 217099, 5273366, 6344769, 7754323, 5835422, 4346988, 7384008, 2772347, 9470450, 8410014, 461561, 1419839, 386641, 10446408, 6583359, 5894522, 6982289, 8025238, 2003728, 7747063, 8751004, 5418343, 10205618, 8414155, 3193310, 2826742, 804294, 8840565, 4905550, 1680516, 1791600, 3234304, 5035622, 3268201, 5113125, 8674985, 5900391, 10552891, 4330548, 5775364, 1215496, 6177987, 59317, 3443628, 239240, 4113298, 7249714, 1283193, 2831705, 4716479, 940438, 8616394, 5245183, 3032536, 8477002, 1690047, 878277, 5946740, 10251949, 2585182, 3213707, 7256810, 8615740, 10497679, 2343030, 7891543, 7648253, 5901695, 6665755, 1767790, 2753149, 5518646, 8559409, 10173671, 2297683, 1220901, 261724, 9399293, 9307070, 267951, 5158111, 9712629, 7491573, 4325073, 6302661, 6864595, 10151497, 10528530, 8727410, 9038033, 7784464, 5018569, 579243, 8107321, 2140312, 7649046, 2551609, 5754089, 5167891, 6729174, 4765625, 9983266, 202105, 1066498, 7204834, 513998, 2151288, 962467, 8683956, 6604701, 1872224, 2475278, 6941401, 3244060, 7071807, 3886125, 1816155, 8099618, 9871810, 1272403, 3776315, 2501920, 10403194, 4082262, 8835519, 1103604, 4943733, 7557491, 8034232, 1847710, 98025, 5817556, 9609957, 3274967, 8233177, 1881204, 5477371, 835608, 1292131, 8590383, 4276184, 193275, 5015759, 87454, 8726848, 4581414, 9764339, 761260, 5019271, 6114818, 2374035, 2674815, 659440, 9657893, 9259479, 3433723, 591767, 4897738, 375661, 724393, 8790787, 95698, 2285328, 5376408, 9969705, 6255985, 754087, 6313606, 3082188, 8415045, 10482332, 10040024, 2337000, 2205257, 3041495, 9425732, 9691941, 7336914, 2751332, 6143264, 7375894, 5313260, 5244324, 10507971, 7874350, 1732929, 52970, 4269062, 8192472, 348120, 6038117, 2457217, 6617007, 7654208, 8063381, 7998295, 10358313, 165594, 5952231, 6230926, 4978381, 6560900, 1540646, 3462570, 9877270, 1494288, 9836255, 1580248, 3469746, 3514845, 1769141, 10046939, 5234865, 38202, 4877146, 737808, 6082095, 8876104, 7423686, 4737138, 6080758, 10121549, 2637734, 8761740, 8929607, 9439686, 1744848, 8740147, 4672532, 5490945, 1440020, 4415326, 1886714, 2760366, 2592368, 518697, 2417745, 4451875, 1954261, 5501000, 605317, 3866060, 1864314, 3571392, 3453831, 8824957, 713853, 10586348, 6482749, 4186030, 4507976, 6808629, 6154628, 4844043, 7694677, 7435451, 3468239, 9136959, 10476342, 1875342, 682777, 8656903, 9396446, 8673442, 8629099, 8783465, 10146652, 10561148, 369183, 8204064, 4482603, 3016378, 1544182, 3002275, 8268624, 176522, 6224534, 1810850, 10298405, 6326790, 9178142, 8059387, 5869821, 5210846, 3624478, 4348241, 9634945, 7194377, 675029, 7198993, 304031, 4130051, 4952602, 10007292, 1632869, 884605, 1846037, 8096151, 9209566, 248981, 1220949, 6603133, 9414910, 1562983, 5863123, 5368399, 10108884, 7960359, 2024638, 355868, 7287271, 55954, 6511388, 7482656, 7786321, 4902137, 10462241, 5325712, 132297, 5371077, 5759762, 9735566, 2880124, 8664526, 9236402, 3776853, 7659355, 8852689, 668998, 3096080, 4680201, 4080899, 5250111, 5470563, 7156111, 772289, 5352938, 3593467, 2285507, 3622954, 9035938, 7488317, 10231034, 9741931, 8389348, 3048716, 7953245, 2076130, 8076098, 472265, 5680315, 8558894, 4209135, 10453504, 3524240, 2141887, 4867051, 138540, 9423043, 7420631, 713808, 7387860, 7995829, 4952214, 4032973, 9773217, 204759, 9729332, 208464, 10310156, 6957457, 2976998, 1019110, 9055248, 1512267, 9522113, 1552753, 9606604, 7907158, 8135589, 8006682, 3471960, 9961023, 8916096, 6951747, 3595419, 2226302, 3364483, 5728061, 3511440, 10263107, 2499707, 3665668, 3488453, 1818772, 1628857, 1763953, 1780995, 5054314, 87513, 4671497, 7915668, 82934, 6686631, 4889800, 306425, 4653307, 1984428, 6512212, 260062, 9208972, 2364915, 1733290, 7656244, 3982331, 2769465, 4717290, 6159826, 2158263, 2528449, 1883341, 8403753, 7179561, 4029280, 4046789, 7417929, 6166801, 944228, 8667000, 8074972, 7031732, 1160175, 9578602, 7270253, 355194, 669222, 462412, 484426, 9176645, 4238037, 3752812, 3793878, 3662983, 521594, 2129629, 7261011, 84682, 1551499, 4871869, 8079028, 4139621, 4302788, 5158021, 581459, 7550315, 8316749, 2823813, 8147036, 10210030, 532804, 7429525, 8700383, 5344708, 9490257, 3672495, 8128214, 9949282, 5309575, 10117133, 5970280, 7025855, 3969793, 7443207, 6467967, 1408301, 443294, 2739543, 3415197, 6364116, 7594158, 5807529, 6785809, 3824253, 3067943, 6451229, 5661488, 3591191, 6799081, 3617345, 6420906, 1060674, 2963285, 8808081, 6498491, 5851815, 2601462, 2160417, 5764832, 5874242, 7719023, 6161292, 8460190, 5990841, 6926834, 4643353, 5468000, 4074631, 5067719, 5293204, 4341230, 10266012, 2579366, 9192406, 6019515, 2902451, 5881065, 9844203, 8898488, 2493956, 9452155, 7366669, 9246293, 394053, 42390, 2048443, 6378340, 2694095, 5096336, 8491970, 1899896, 4538275, 3495235, 4544994, 3101458, 6401365, 453372, 6679104, 7963796, 3779428, 6787221, 9010171, 2110958, 8551613, 10453706, 1001130, 1290456, 4460761, 2101379, 8110367, 4204146, 4418491, 2157550, 5389013, 5609170, 3791526, 479664, 4574094, 6975177, 648044, 1095420, 10585045, 2537872, 2682216, 6476887, 6771882, 3592092, 4197570, 10096997, 2631697, 5116270, 5324427, 3959215, 272165, 1840267, 4614002, 4549128, 6667635, 7805257, 4709944, 8335458, 9219843, 7180314, 9310453, 9277750, 1185864, 9410397, 4466681, 8413578, 5765171, 4426002, 6008348, 2940052, 5075432, 10305601, 2876086, 7526934, 7965318, 8710680, 7229725, 5183871, 122842, 722502, 2183837, 5718599, 4749046, 9596738, 652772, 2369815, 6787247, 6862497, 1690139, 3892500, 426270, 3173366, 915627, 3486815, 2870877, 3895647, 3508451, 133056, 3249183, 4757975, 7169863, 1777822, 6467054, 3923863, 4730888, 9146195, 6302745, 3850328, 633060, 3190201, 2234332, 8347861, 3097890, 1889136, 1609440, 5745710, 9116210, 8605173, 3805762, 7550018, 3032484, 1888226, 4880260, 1089341, 6004825, 8280547, 4722136, 3787441, 1690236, 8692893, 4095313, 3035009, 7804730, 4558499, 4158470, 4257476, 5835816, 866801, 3976150, 7765959, 4099919, 3076366, 3407259, 3076366, 8653065, 4794161, 10099073, 2135637, 820131, 3787381, 2534891, 7234456, 520298, 3092074, 2897989, 6201828, 4915708, 8877384, 3161376, 5008800, 1522999, 10295695, 3724681, 4246570, 3607931, 10196541, 8730525, 404839, 2126804, 10551850, 8400851, 8793025, 3451000, 2862075, 4439106, 9236156, 501246, 5615318, 8797876, 1970962, 4791572, 5626428, 2492978, 1126118, 3381989, 8501276, 10403865, 8508958, 10197346, 7567004, 3061037, 9742506, 6653770, 3939071, 10574914, 7810535, 750625, 7524795, 8802641, 1993500, 6687292, 9737738, 6848, 10095329, 1737468, 1292191, 4207174, 4226094, 2725570, 8733860, 2816637, 3211172, 7162259, 8920561, 7034054, 4369432, 9295897, 6301479, 9225378, 10214246, 1308529, 6108296, 7332486, 1512968, 2967106, 1153227, 1821575, 8374617, 9946222, 5098466, 6307863, 8985643, 7170482, 2242309, 6782305, 6556822, 2922392, 4977078, 10059467, 7413126, 8671422, 2800019, 5760063, 6706157, 9642795, 7467980, 7840849, 10186063, 307352, 9040587, 3482865, 4698239, 1450583, 8756110, 2820494, 2653419, 214702, 3142601, 170072, 6570772, 8551242, 9582038, 6074828, 8918041, 5509838, 3978869, 9999412, 4057580, 8630070, 10147538, 8229850, 6175896, 785759, 6462251, 1451359, 1609219, 8649896, 8361502, 97794, 3359242, 7280041, 9553133, 5614987, 7115544, 7128667, 3531161, 1523289, 10444322, 210449, 6245796, 9513184, 9265778, 809724, 6997280, 2255638, 8927387, 451632, 7858532, 2640316, 1690436, 43205, 7973795, 9848865, 1479652, 8137174, 7830541, 283202, 2386530, 2142818, 6461058, 7498013, 7949912, 7477393, 8355285, 1125886, 350220, 3024946, 7422107, 2347057, 2119449, 8502146, 5654146, 2678754, 7396723, 4929381, 5991298, 5773493, 8298256, 9276416, 3432862, 2401180, 10359631, 6693088, 9851805, 7413803, 5135171, 4432545, 9921924, 3944877, 2397021, 836446, 2459385, 5303768, 3660771, 3248133, 2008648, 5949197, 9017328, 233492, 1249249, 7681691, 5444902, 9216214, 3785741, 5760408, 3922970, 1690102, 6702165, 5975904, 3423380, 2923374, 6927084, 6396621, 93560, 7952506, 9747047, 2279324, 9025895, 6901005, 8933429, 4529957, 7464242, 3756477, 2058409, 9044583, 9109507, 9598926, 2290838, 4458625, 8297872, 10090252, 7376648, 5085175, 5662671, 6542510, 4634326, 4586608, 3673592, 6060537, 4253559, 5421105, 10410098, 2782512, 3853326, 7758322, 7904285, 2345963, 3460601, 5537657, 441720, 5208084, 4818751, 101477, 5673796, 7023589, 4674536, 5835339, 7886212, 2238189, 7644232, 9837157, 10253479, 3748748, 8237151, 3028989, 1700649, 7466537, 10342763, 1394706, 4702582, 7835068, 4166433, 10478601, 7842814, 2590035, 4778912, 4206959, 6411132, 7704401, 4460683, 7002935, 7808104, 6417678, 8288692, 2442355, 4750492, 5278083, 6518857, 6149987, 8984118, 8892414, 7880797, 7043021, 9797118, 4289256, 1560609, 1592810, 5981122, 5668603, 9733777, 8930228, 823476, 10395235, 5151590, 1266105, 1409519, 5438671, 9437264, 1393232, 1867247, 1198472, 5811002, 3577641, 605154, 1639381, 6848897, 1949783, 361831, 3573839, 3609126, 5897132, 9586866, 4201796, 2852238, 3848730, 2976893, 10290446, 3231558, 7642775, 6778950, 10099356, 87797, 9735197, 6533162, 9289863, 1901216, 4288047, 5710832, 2931582, 5644407, 8734536, 8663419, 7377451, 7174863, 1323094, 10279294, 4455777, 1287110, 10296967, 9100528, 8787444, 239585, 4016710, 954010, 6483481, 10434928, 2202949, 9075864, 8757088, 1858036, 679769, 1753809, 10263060, 5358284, 898998, 7452691, 738085, 1249587, 584971, 6298118, 1703736, 6178539, 2575996, 6108439, 9443298, 7168472, 1859991, 8306042, 4570472, 10485237, 162903, 1782067, 8805251, 9062469, 1474664, 3369446, 824766, 3717110, 9613788, 165937, 4530916, 6878348, 2056604, 9680495, 1773697, 5066657, 5910628, 10418530, 4872006, 3822365, 1665476, 5885552, 8265329, 7091740, 393576, 10340313, 4025480, 6998200, 5875808, 3097416, 6569684, 5057406, 1263212, 10512132, 756006, 3077846, 4752772, 2036451, 8395730, 1293683, 2281345, 789381, 3640957, 4314276, 1934541, 3807256, 7197157, 9479883, 4110234, 9795073, 7818258, 832531, 8411271, 4667733, 3528138, 2885225, 4498853, 3842012, 9292773, 546952, 8426036, 7938618, 388318, 2469982, 9150950, 9741685, 7542995, 3281385, 8206910, 3691910, 9931837, 4143058, 9653536, 10357949, 4465042, 3740035, 8076586, 5449842, 10348590, 10529707, 3790443, 2157296, 6461298, 3811084, 2056890, 27886, 3442656, 3272794, 243522, 74558, 9649019, 7123976, 4513895, 3676616, 9597241, 9777177, 2950710, 7031836, 3713676, 4994244, 7918306, 1703223, 850420, 770193, 839068, 7811, 8947602, 4687618, 6548721, 2575573, 6278536, 9129860, 5306813, 4962075, 6307418, 3668488, 3765785, 7492284, 9263267, 9013404, 9894766, 3502421, 3720735, 7869708, 7665187, 9195323, 6974825, 5652987, 3758086, 1313608, 580177, 6393212, 678009, 6511472, 8864792, 3798632, 6127761, 353156, 5629720, 5168905, 5276869, 6798449, 1043082, 1835418, 5178938, 277839, 213002, 912364, 7734679, 8519839, 9376989, 7794463, 8537232, 6236267, 8035163, 2008957, 7863748, 6060840, 6649959, 10142988, 10297036, 1856275, 6562458, 105291, 10444224, 137643, 854886, 1746338, 6834789, 978801, 1804869, 2259893, 275334, 6827205, 3230151, 649080, 2105104, 9514911, 663852, 8816664, 4183316, 10191905, 550998, 7923488, 4403855, 1285445, 7038109, 9509130, 3920334, 8743417, 5896889, 1197215, 3735513, 8369786, 8995327, 3250095, 7904698, 10163600, 4387316, 5904322, 601717, 10559480, 4543486, 2209693, 3797622, 477909, 5807894, 7392240, 8195685, 8500980, 8568292, 3183507, 10501131, 9205975, 5251443, 8982699, 2000178, 6378331, 9035866, 8061417, 6781998, 2924362, 3348655, 878595, 9211115, 7883722, 9712962, 7872754, 1725440, 9098559, 8898486, 2485198, 8179017, 9829021, 9896640, 9987071, 7512659, 9001608, 1937517, 6824434, 10073509, 6821629, 4054019, 2124778, 3607195, 4647904, 4074299, 4184019, 2141966, 1182507, 238682, 9288602, 10116820, 1650195, 2832902, 10385820, 6239044, 2326231, 1293822, 1749238, 1087001, 2278103, 5105112, 9189609, 9556196, 2965288, 5677643, 8044589, 8599126, 5255854, 2305113, 2149817, 5522106, 8777048, 868318, 3901210, 8485108, 3861708, 6393131, 2370429, 9781886, 7954231, 2362868, 2373907, 3605564, 9637521, 5901459, 4179861, 8811923, 6120456, 3732433, 5206313, 6145698, 7662111, 4979899, 1896842, 895166, 9279070, 9282397, 4090746, 4790110, 7578779, 6102804, 8452208, 362300, 1176787, 7693516, 826760, 854681, 8299242, 2357412, 10154273, 379972, 6420895, 978908, 3178439, 8032263, 1127637, 8204191, 4567152, 7214745, 4187014, 9999127, 9692174, 3172113, 7982496, 2223958, 80197, 7780691, 5512922, 9553545, 9528008, 1408293, 5295883, 9974577, 4946641, 6603567, 9213237, 4826557, 1347639, 1353241, 10105079, 1284891, 8740786, 6816495, 7542239, 2873288, 8005824, 8738136, 9393396, 1481094, 4921562, 2817359, 2780135, 5540945, 3009661, 4335335, 7713942, 10166477, 3565448, 1820016, 1759631, 9963604, 4504187, 138888, 124395, 4516458, 8343667, 8860339, 4341599, 2579388, 10086152, 566051, 9210292, 3106632, 7346002, 5738417, 1645529, 1529658, 8730092, 5430712, 8030289, 5341445, 9974441, 3780589, 3182136, 1662852, 2430893, 78189, 8685955, 3004140, 8886778, 9541545, 6584503, 2476591, 4144989, 8590358, 4833301, 6346826, 7706445, 928976, 8426810, 6765924, 10035578, 6513027, 9121141, 5941526, 9585981, 3237337, 6881315, 2632303, 6019273, 2453321, 4825904, 3375804, 5954960, 6788980, 8852897, 5854923, 1877025, 3273377, 1840519, 1286592, 2330940, 10121584, 4072853, 7387511, 2504381, 9030212, 2048915, 427993, 3895114, 3471873, 9893632, 484468, 1731890, 8233010, 9096639, 4122047, 4140448, 834186}
	FinlandNodes = []int{728652, 110611, 626130, 715935, 731568, 545875, 239476, 111993, 22456, 484102, 207355, 813582, 704462, 800619, 319546, 447197, 433927, 210866, 397484, 191488, 752941, 858119, 225867, 558678, 751816, 873106, 606328, 44716, 545995, 191895, 417374, 581142, 5765, 766736, 441376, 754878, 137704, 110343, 877904, 462092, 703282, 321789, 271075, 74256, 708213, 487186, 864125, 434782, 897834, 668585, 481143, 328200, 755791, 787787, 804898, 517329, 38329, 385557, 52910, 147543, 332693, 508533, 78022, 182641, 755700, 588135, 171704, 792301, 404424, 12331, 463707, 682149, 514431, 660553, 50353, 301574, 111580, 159890, 697437, 433239, 880503, 213494, 821491, 434268, 463123, 185742, 74775, 267689, 234777, 882035, 629121, 124414, 805268, 34744, 238945, 426994, 358980, 467814, 882226, 348105, 49877, 421225, 886172, 660210, 808690, 117593, 552176, 353276, 381605, 479289, 143576, 215157, 459668, 107667, 297354, 34686, 321908, 485800, 230598, 67389, 872273, 221739, 732023, 587348, 753484, 406863, 418034, 24114, 658888, 585920, 594738, 594682, 9963, 399355, 105233, 235726, 546929, 186023, 670348, 609799, 783980, 513459, 773869, 480130, 108777, 92163, 175665, 828870, 722232, 69790, 790840, 803213, 151523, 828589, 376782, 134263, 253693, 790789, 152821, 365474, 232532, 602094, 835697, 211793, 510894, 764759, 295461, 245014, 598198, 155135, 512634, 483572, 279703, 740451, 83830, 29656, 582624, 57196, 294203, 648408, 67383, 628856, 435138, 418317, 558189, 43173, 651859, 412915, 239282, 502303, 522995, 484783, 322710, 498161, 402752, 536297, 20764, 765584, 449088, 276625, 17355, 533171, 834635, 254981, 383510, 209309, 891922, 772809, 477182, 367671, 667245, 367083, 337482, 725056, 902133, 852205, 401564, 140367, 503914, 131911, 212593, 69331, 181539, 168629, 283736, 183020, 295266, 548967, 518341, 707482, 125622, 434149, 398782, 741374, 8138, 419563, 268158, 618699, 503931, 528594, 37995, 490222, 764459, 407982, 687149, 91850, 890969, 668736, 143210, 480549, 698890, 12724, 803503, 277601, 355647, 257762, 261803, 424876, 423287, 452671, 212446, 406612, 433707, 282632, 174221, 675019, 268424, 192268, 768974, 615963, 52, 180939, 774339, 271686, 77732, 37234, 731231, 713928, 109461, 735758, 253337, 197115, 333738, 510371, 569408, 816725, 37155, 807837, 46437, 746214, 121659, 698470, 43366, 269891, 52628, 12462, 401050, 660123, 899435, 755973, 822040, 654775, 295733, 184192, 185815}
)

func BenchmarkPointCorrection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CorrectPoint(server.Config, Coord{62.24, 25.74}, "finland")
	}
}

func get_country_coords(nodes []int, country string) error {
	_, err := GetCoordinates(server.Config, country, nodes)
	return err
}

func BenchmarkFinlandPointToCoordinateTranslation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_country_coords(FinlandNodes, "finland")
	}
}

func BenchmarkGermanyPointToCoordinateTranslation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_country_coords(GermanyNodes, "germany")
	}
}

func TestGermany(t *testing.T) {
	query := `
	EXPLAIN ANALYZE
SELECT n.coord[0], n.coord[1]
FROM (
	SELECT *, arr[rn] AS node_id
	FROM (
		SELECT *, generate_subscripts(arr, 1) AS rn
		FROM (
			SELECT ARRAY%s AS arr
		) x
	) y
) z
JOIN germany_nodes n ON n.id = z.node_id
ORDER BY z.arr, z.rn;`
	values := strings.Replace(fmt.Sprintf("%v", GermanyNodes), " ", ",", -1)
	q := fmt.Sprintf(query, values)

	db, _ := sql.Open("postgres", server.Config.String())
	defer db.Close()

	rows, err := db.Query(q)

	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var msg string
		if err := rows.Scan(&msg); err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	}
}

func BenchmarkAPIPointCorrection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query_string := fmt.Sprintf("lat=%f&long=%f&country=%s", 62.24, 25.74, "finland")
		request := fmt.Sprintf("http://localhost:%d/point?%s", server.Config.Port, query_string)
		if _, err := http.Get(request); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAPIPointsToCoordinates(b *testing.B) {
	var countries = []string{"finland", "germany"}
	for country := range countries {
		for i := 0; i < b.N; i++ {
			values := strings.Replace(fmt.Sprintf("%v", GermanyNodes), " ", ",", -1)
			query_string := fmt.Sprintf("nodes=%s&country=%s", values, country)
			request := fmt.Sprintf("http://localhost:%d/coords?%s", server.Config.Port, query_string)
			if _, err := http.Get(request); err != nil {
				b.Fatal(err)
			}
		}
	}
}
