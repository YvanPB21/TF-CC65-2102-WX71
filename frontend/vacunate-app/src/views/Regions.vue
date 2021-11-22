<template>
  <v-container style="padding-bottom: 100px;padding-left: 20%;padding-right: 20%;">
    <div align="center">
      <v-row class="subtitle">
        <p>Predicciones por región</p>
      </v-row>
    </div>
    <v-expansion-panels>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://www.peru.travel/Contenido/Destino/Imagen/pe/37/1.3/Principal/Los%20Uros.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('PUNO')">
          PUNO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://www.esan.edu.pe/conexion/actualidad/2018/10/16/1500x844_lima_2035_2050.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('LIMA')">
          LIMA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('LAMBAYEQUE')" style="color: white;font-weight: bolder ;background-image: url('https://sobrehistoria.com//wp-content/uploads/2008/02/el-se%C3%B1or-de-sipan-la-cultura-mochica.jpg')">
          LAMBAYEQUE
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('ANCASH')" style="color: white;font-weight: bolder ;background-image: url('https://www.peru.travel/Contenido/Destino/Imagen/es/2/1.3/Principal/Ancash.jpg')">
          ANCASH
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://astelus.com/wp-content/viajes/arequipa-y-el-volcan-misti-peru.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('AREQUIPA')">
          AREQUIPA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('LORETO')" style="color: white;font-weight: bolder ;background-image: url('https://portal.andina.pe/EDPfotografia3/Thumbnail/2017/11/01/000460646W.jpg')">
          LORETO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://www.peru.travel/Contenido/General/Imagen/es/474/1.1/Portada_Danzante_Tijeras.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('HUANCAVELICA')">
          HUANCAVELICA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('JUNIN')" style="color: white;font-weight: bolder ;background-image: url('https://blog.redbus.pe/wp-content/uploads/2019/10/2019-10-04.png')">
          JUNIN
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: white;font-weight: bolder ;background-image: url('https://thumbs4.properati.com/4/zZ0Wqw_gGVcH7_Ex2x0WkSxk6nY=/trim/fit-in/0x0/smart/filters:strip_icc()/https://properati-s3-pe.s3-us-west-2.amazonaws.com/temp-developments-editor-67782bdc-8661-4a3a-a891-9a256e76f46d-GettyImages-1014260912.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('PIURA')">
          PIURA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('CAJAMARCA')" style="color: black;font-weight: bolder ;background-image: url('https://siar.regioncajamarca.gob.pe/sites/default/files/archivos/public/docs/lecheras.jpg')">
          CAJAMARCA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('AYACUCHO')" style="color: black;font-weight: bolder ;background-image: url('https://content.r9cdn.net/rimg/dimg/21/59/4f44151c-city-40948-169fd26ee9a.jpg?width=1200&height=630&crop=true')">
          AYACUCHO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('UCAYALI')" style="color: black;font-weight: bolder ;background-image: url('https://www.ipe.org.pe/portal/wp-content/uploads/2010/10/Ucayali-870x500.jpg')">
          UCAYALI
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header  v-on:click="predictRegion('TUMBES')" style="color: black;font-weight: bolder ;background-image: url('https://www.go2peru.com/destinos/large/Tumbes_Feb08-699.jpg')">
          TUMBES
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://peruconstruye.net/wp-content/uploads/2020/09/callao.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('CALLAO')">
          CALLAO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('LA LIBERTAD')"  style="color: white;font-weight: bolder ;background-image: url('https://imgs.mongabay.com/wp-content/uploads/sites/25/2017/02/09155629/bosques_secos-deforestacion-conservacion-biodiversidad-peru-7.jpg')">
          LA LIBERTAD
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://content.emarket.pe/common/collections/standard/6f/26/6f263b90-edcd-4662-96b7-7dd0e47d6677.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('SAN MARTIN')">
          SAN MARTIN
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://propuestapais.pe/wp-content/uploads/2019/10/Ica.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('ICA')">
          ICA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://d1bvpoagx8hqbg.cloudfront.net/originals/apurimac-alturas-canones-naturales-fd503277fa8c8610123683d60f483e07.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('APURIMAC')">
          APURIMAC
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://incarail.com/blog/wp-content/uploads/2019/03/TEMPORADA-SECA1.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('CUSCO')">
          CUSCO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('https://cdn.www.gob.pe/uploads/document/file/2011127/standard_0005-min.jpg.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('HUANUCO')">
          HUANUCO
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('AMAZONAS')" style="color: black;font-weight: bolder ;background-image: url('https://www.passporthealthglobal.com/wp-content/uploads/2019/05/CONOCE-EL-AMAZONAS.jpg')">
          AMAZONAS
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header v-on:click="predictRegion('MADRE DE DIOS')" style="color: black;font-weight: bolder ;background-image: url('https://www.ytuqueplanes.com/imagenes/fotos/regiones/RC-BANNER-INT-Madre-de-Dios.jpg')">
          MADRE DE DIOS
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: white;font-weight: bolder ;background-image: url('https://portal.andina.pe/EDPfotografia/Thumbnail/2014/01/24/000231459W.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('TACNA')">
          TACNA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: white;font-weight: bolder ;background-image: url('https://www.peru.travel/Contenido/Evento/Imagen/pe/103/1.1/Evento/aniversario-moquegua.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('MOQUEGUA')">
          MOQUEGUA
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel
              style="color: black;font-weight: bolder ;background-image: url('http://arellanojuan.com/wp-content/uploads/2015/10/cdp.jpg')"
      >
        <v-expansion-panel-header v-on:click="predictRegion('PASCO')">
          PASCO
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
      regions:[],
      selected:undefined,
    }
  },
  methods:{
    predictRegion(region){
      //axios.get("http://localhost:8080/prediction?dep="+region)
    console.log(region)
    }

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
