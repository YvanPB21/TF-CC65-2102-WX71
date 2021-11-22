<template>
  <div >
    <v-row >
      <v-col class="main" >
        <v-row class="top" >
          <v-img
            src="../assets/logov.png"
            max-height="60"
            contain="true"
            style="position: relative; right: 200px"

          ></v-img>
        </v-row>
        <v-row style="padding: 0;margin-bottom: 0;">
          <v-img src="../assets/banner.png" position="center" width="800px" height="600px" contain="true"></v-img>
        </v-row>
        <div  align="center">
          <router-link to="/Regions" style="text-decoration: none">
            <v-btn
                    color="#32ED3A"
                    elevation="5"

            >
              <span style="color: white; font-weight: bolder">Ver predicciones</span>
            </v-btn>
          </router-link>
        </div>
        <div  style="padding-bottom: 50px">
          <v-row class="subtitle">
            <p>Medidas adoptadas para la prevención del COVID-19</p>
          </v-row>
          <div class="recomendaciones" style="padding-bottom: 25px">
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

        <div style="background-color: #F4F4F7;padding-bottom: 25px">
          <v-row class="subtitle">
            <p>Actualidad del COVID-19 en el Perú (diario)</p>
          </v-row>
          <div class="line">
            <div style="background-color: #F5D6D4; padding: 15px">
              <div style="display: flex" class="rojo">Casos confirmados</div>
              <div class="rojo">962 (5 más que ayer)</div>
            </div>
            <div style="background-color: #D4F5D9; padding: 15px">
              <div style="display: flex" class="verde">Altas hospitalarias</div>
              <div class="verde">110 (27 más que ayer)</div>
            </div>
            <div style="padding: 15px; background-color: #DAD4D4">
              <div style="display: flex" class="negro">Fallecidos</div>
              <div class="negro">13 (4 menos que ayer)</div>
            </div>
          </div>
          <div  align="center">
            <v-btn
                    color="red"
                    elevation="5"
                    href="https://www.gob.pe/coronavirus"
            >
              <span style="color: white; font-weight: bolder">Más información</span>
            </v-btn>
          </div>
        </div>
        </div>
        <div align="center">
          <v-row class="subtitle">
            <p>Centros de vacunación</p>
          </v-row>
          <div align="center">
            <img src="https://image.flaticon.com/icons/png/512/1097/1097213.png" width="80px" />
          </div>
          <span>Los centros de vacunación están por toda la ciudad, averigue cuál es el más cercano a usted.</span>
          <div style="padding: 10px;margin-bottom: 80px;margin-top: 20px" align="center">
            <v-btn
                    color="red"
                    elevation="5"
                    href="https://www.gob.pe/13334"
            >
              <span style="color: white; font-weight: bolder">Buscar centros</span>
            </v-btn>
          </div>
        </div>

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
  margin-top: 50px;
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
.line{
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  text-align: center;
  padding-bottom: 40px;
}
.rojo{
  color: #B10C0C;
  font-weight: bolder;
}
.verde{
  color: #106D24;
  font-weight: bolder;
}
.negro{
  color: #242A25;
  font-weight: bolder;
}
.main{
  margin-right: 150px;
  margin-left: 150px;
  
}
</style>
