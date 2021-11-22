<template>
  <v-container style="padding-bottom: 100px">
    <div align="center">
      <v-row class="subtitle">
        <p>Predicciones por región</p>
      </v-row>
    </div>
    <v-expansion-panels>
      <v-expansion-panel
        v-for="(item,i) in regions"
        :key="i"
        style="background-image:url('https://www.esan.edu.pe/conexion/actualidad/2018/10/16/1500x844_lima_2035_2050.jpg')"
      >
        <v-expansion-panel-header>
          {{item}}
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-container>
</template>

<script>
  //import axios from "axios"
export default {
  name: 'Regions',
  data(){
    return{
      regions:["PUNO","LIMA","LAMBAYEQUE","ANCASH","AREQUIPA","LORETO","HUANCAVELICA","JUNIN","PIURA","CAJAMARCA",
      "AYACUCHO","UCAYALI","TUMBES","CALLAO","LA LIBERTAD","SAN MARTIN","ICA","APURIMAC","CUSCO","HUANUCO","AMAZONAS",
      "MADRE DE DIOS","TACNA","MOQUEGUA","PASCO"],
      selected:undefined,
    }
  },
  methods:{

  },
  created() {
  }
}
</script>
<style>
  .subtitle {
    text-align: center;
    justify-content: center;
  }

  .subtitle p {
    font-size: 2em;
    font-family: "Trebuchet MS", "Lucida Sans Unicode", "Lucida Grande",
    "Lucida Sans", Arial, sans-serif;
    margin: 40px;
  }
</style>
