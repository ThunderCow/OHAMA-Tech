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
Si at det binære tallet er 0010 1010 1111&#8322;, da kan vi se i tabellen at 0010=2, 1010=A og 1111=F. 0010 1010 1111=0x2AF. Det samme prinsippet gjelder dersom det skal konverteres fra heksadesimal til binære tall, bare motsatt.
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

For å konvertere binære tall om til desimaler setter vi opp det binære tallet i en tabell. Si at det binære tallet er som i eksempelet over, 0010 1010 1111.
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
512+128+32+8+4+2+1=687. det binære tallet 0010 1010 1111 er da 687 i desimalform. For å gjøre om desimal til et binært tall tar vi 687^10=0000 0010 1010 1111^2.
<br>
