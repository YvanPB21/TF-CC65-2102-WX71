<template>
  <div class="home">
    <v-row >
      
      <v-col class="main">
        <v-row class="top">
          <v-img
            src="../assets/logov.png"
            max-height="60"
            contain="true"
          ></v-img>
        </v-row>
        <v-row>
          <v-img src="../assets/banner.png" max-width="800px" contain="true"></v-img>
        </v-row>

        <v-row class="subtitle">
          <p>Medidas adoptadas para la prevención del COVID-19</p>
        </v-row>
        <div class="recomendaciones">
          <v-col>
            <img
              src="https://cdn-icons-png.flaticon.com/512/3579/3579748.png"
              width="80px"
            />
            <p>Todas las personas de 2 años de edad o más que no estén totalmente vacunadas deben usar una mascarilla en lugares públicos cerrados.</p>
          </v-col>
          <v-col>
            <img src="../assets/hand.png" width="80px" />
            <p>Lávese las manos con frecuencia con agua y jabón por al menos 20 segundos, especialmente después de haber estado en un lugar público, o después de sonarse la nariz, toser o estornudar.</p>
          </v-col>
          <v-col>
            <img src="../assets/vacc.png" width="80px" />
            <p>Las vacunas autorizadas contra el COVID-19 pueden ayudar a protegerlo del COVID-19.</p>
          </v-col>
        </div>
        <v-row></v-row>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "Home",
  components: {},
};
</script>

<style scoped>
.top {
  margin-top: 10px;
  margin-bottom: 10px;
}

.recomendaciones {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  text-align: center;
}
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

.main{
  margin: 0px 200px 0px 200px;
}
</style>
