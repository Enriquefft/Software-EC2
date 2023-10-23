<template>
  <LocationSelector @get_distance="show_distance" />
  <DistanceDisplay :city1="city1" :city2="city2" />

  <!--city1-->
  <span>City 1: </span>
  <span>{{ city1 }}</span>
</template>

<script>
import LocationSelector from "./components/LocationSelector.vue";
import DistanceDisplay from "./components/DistanceDisplay.vue";

export default {
  name: "App",
  components: {
    LocationSelector,
    DistanceDisplay,
  },
  data() {
    return {
      city1: "",
      city2: "",
    };
  },

  methods: {
    show_distance(city1, city2) {
      // console.log(event);

      this.city1 = city1;
      this.city2 = city2;

      console.log(city1, city2);
      console.log(this.city1, this.city2);
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
