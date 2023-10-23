<template>
  <div>
    <p class="text-center">Distancia entre ciudades</p>
    <p class="text-center">Ciudad 1: {{ city1 }}</p>
    <p class="text-center">Ciudad 2: {{ city2 }}</p>
  </div>
</template>

<script>
export default {
  props: ["city1", "city2"],
};
</script>

<!-- <script> -->
<!-- export default { -->
<!--   data() { -->
<!--     return { -->
<!--       ciudad1: "", -->
<!--       ciudad2: "", -->
<!--     }; -->
<!--   }, -->
<!--   methods: { -->
<!--     async enviarSolicitud(tecnica) { -->
<!--       const url = `http://192.168.19.7:3001/distance?city1=${this.ciudad1.toLowerCase()}&city2=${this.ciudad2.toLowerCase()}&method=${tecnica}`; -->
<!--       console.log(url); -->
<!---->
<!--       try { -->
<!--         const response = await this.$axios.post(url); -->
<!--         console.log(response); -->
<!--         this.respuesta = response.data; -->
<!--         console.log(response.data); -->
<!--       } catch (error) { -->
<!--         console.error("Error al hacer la solicitud:", error); -->
<!--         console.log(this.respuesta); -->
<!--         this.respuesta = "Error al hacer la solicitud."; -->
<!--       } -->
<!--     }, -->
<!--   }, -->
<!-- }; -->
<!-- </script> -->
