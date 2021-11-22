<template>
  <v-app>
    <v-main>
      <router-view></router-view>
      <footer style="background-color: black; text-align: center;color: white; font-weight: bolder;padding: 10px;position: absolute;bottom: 0; width: 100%">
        <span>Vacúnate Contra Todo Mal 2021</span>
      </footer>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  components: {
  },

  data: () => ({
    //
  }),
};
</script>
