<table style="width:100%">
  <tr>
    <th>Binære tall</th>
    <th>Heksadesimale tall</th> 
    <th>Desimaltall</th>
  </tr>
  <tr>
    <td>1101</td>
    <td>0xD</td> 
    <td>13</td>
  </tr>
  <tr>
    <td>1101 1110 1010</td>
    <td>0xDEA</td> 
    <td>3562</td>
  </tr>
  <tr>
    <td>1111 0011 0100</td>	 	
    <td>0xAF34</td>
    <td>44852</td>
  </tr>
  <tr>
    <td>1111 1111 1111 1111</td>
    <td>0xFFFF</td>
    <td>65535</td>
  </tr>
  <tr>
    <td>0001 0001 0111 1000 1010</td>
    <td>0x1178A</td>
    <td>71562</td>
  </tr>
</table>
<br>

<b>A)</b>
For å gjøre om de binære tallene om til heksadesimaler, kan man følge tabellene for de to tallsystemene. Begge systemene har 16 siffer og man kan enkelt finne ut hvilke verdier som er like dersom man har tabellene satt opp ved siden av hverandre. For eksempel:
Si at det binære tallet er 0010 1010 1111&#8322;, da kan vi se i tabellen at 0010&#8322;=2&#8321;&#8326;, 1010&#8322;=A&#8321;&#8326; og 1111&#8322;=F&#8321;&#8326;. 0010 1010 1111&#8322;=0x2AF&#8321;&#8326;. Det samme prinsippet gjelder dersom det skal konverteres fra heksadesimal til binære tall, bare motsatt.
<br>
<table style="width:100%">
  <tr>
    <th>Binære tall</th>
    <th>Heksadesimale tall</th> 
    <th>Desimaltall</th>
  </tr>
  <tr>
    <td>1101</td>
    <td>0xD</td> 
    <td>13</td>
  </tr>
  <tr>
    <td>1101 1110 1010</td>
    <td>0xDEA</td> 
    <td>3562</td>
  </tr>
  <tr>
    <td>1111 0011 0100</td>	 	
    <td>0xAF34</td>
    <td>44852</td>
  </tr>
  <tr>
    <td>1111 1111 1111 1111</td>
    <td>0xFFFF</td>
    <td>65535</td>
  </tr>
  <tr>
    <td>0001 0001 0111 1000 1010</td>
    <td>0x1178A</td>
    <td>71562</td>
  </tr>
</table>
<br>

For å konvertere binære tall om til desimaler setter vi opp det binære tallet i en tabell. Si at det binære tallet er som i eksempelet over, 0010 1010 1111&#8322;.
<br>
<table style="width:100%">
  <tr>
    <th>0</th>
    <th>0</th>
    <th>1</th>
    <th>0</th>
    <th>1</th>
    <th>0</th>
    <th>1</th>
    <th>0</th>
    <th>1</th>
    <th>1</th>
    <th>1</th>
    <th>1</th>
  </tr>
  <tr>
    <td>2048</td>
    <td>1024</td>
    <td>512</td>
    <td>256</td>
    <td>128</td>
    <td>64</td>
    <td>32</td>
    <td>16</td>
    <td>8</td>
    <td>4</td>
    <td>2</td>
    <td>1</td>
   </tr>
</table>

<br>
Etter vi har satt tallene i en tabell, skal vi addere tallene som befinner seg under tallet 1. 
512+128+32+8+4+2+1=687&#8321;&#8320;. det binære tallet 0010 1010 1111&#8322; er da 687&#8321;&#8320; i desimalform. For å gjøre om desimal til et binært tall tar vi 687&#8321;&#8320;=0000 0010 1010 1111&#8322;.
<br>
<table style="width:100%">
  <tr>
    <th>Binære</th>
    <th>Heksadesimal</th>
  </tr>
  <tr>    
    <td>0000</td>
    <td>2</td>

</tr>  
  <tr>
    <td>0010</td>
    <td>2</td>
  </tr>
  <tr>
    <td>1010</td>
    <td>A</td>
  </tr>
  <tr>
    <td>1111</td>
    <td>F</td>
  </tr>
</table>
<br>

<b>B)</b> Når vi går fra heksadesimaler til desimaler bruker vi tabellene fra hvert tallsystem til å finne verdien sifrene har når vi skal konvertere de. Vi setter opp en tabell for å vise utregningen. Vi bruker 4FBA som et eksempel.
<table style="width:100%">
  <tr>
    <th>Posisjon</th>
    <th>Siffer</th> 
    <th>Utregning</th>
  </tr>
  <tr>
    <td>0</td>
    <td>A</td> 
    <td>10*16<sup>0</sup>=10</td>
  </tr>
  <tr>
    <td>1</td>
    <td>B</td> 
    <td>11*16<sup>1</sup>=176</td>
  </tr>
  <tr>
    <td>2</td>	 	
    <td>F</td>
    <td>15*16<sup>2</sup>=3840</td>
  </tr>
  <tr>
    <td>3</td>
    <td>4</td>
    <td>4*16<sup>3</sup>=16384</td>
  </tr>
 </table>
<br>

Deretter adderer vi svarene vi fikk. 10+176+3840+16384=20410. Når vi konverterer det heksadesimale 4FBA om til desimaltall, får vi 20410. 
Dersom vi skal konvertere fra desimaltall til heksadesimal setter vi tallene i en tabell for å få en god oversikt over regnestykket. Tallene skal alltid divideres på 16, og svaret man får skal multipliseres med 16. Noen av tallene vil ha desimaler bak komma, men de kan man se bort ifra og bare regne videre. Tallet vi får når vi dividerer skal subtraheres med tallet som multipliseres. Da vil vi finne det vi kaller for «rest».  Her er et eksempel på hvordan vi regner fra desimal til heksadesimal. Vi bruker tallet 435.
<br>
<table style="width:100%">
  <tr>
    <th>Verdi</th>
    <th>Utregning</th> 
    <th>Rest</th>
    <th>Heksadesimal rest</th>
  </tr>
  <tr>
    <td>435</td>
    <td>435:16=27</td> 
    <td>3</td>
    <td>3</td>
  </tr>
  <tr>
    <td>27</td>
    <td>27:16=1</td> 
    <td>11</td>
    <td>b</td>
  </tr>
  <tr>
    <td>1</td>	 	
    <td>1:16=0</td>
    <td>1</td>
    <td>1</td>
  </tr>
  <tr>
    <td>0</td>
    <td></td>
    <td></td>
    <td></td>
  </tr>
 </table>
<br>

435:16=27, 27x16=432, 435-432=3
27:16=1, 1x16=16, 27-16=11
1:16=0, 0x16=0, 1-0=1
0
Det heksadesimale tallet blir da 1b3<sap>16</sap>.
