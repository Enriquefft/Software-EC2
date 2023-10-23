<!-- <template> -->
<!--   <GMapAutocomplete placeholder="Ciudad 1" @place_changed="setPlace1"> -->
<!--   </GMapAutocomplete> -->
<!--   <GMapAutocomplete placeholder="Ciudad 2" @place_changed="setPlace2"> -->
<!--   </GMapAutocomplete> -->
<!---->
<!--   {{ ciudad1 }} -->
<!--   {{ ciudad2 }} -->
<!--   <button @click="$emit('get_distance', this.ciudad1, this.ciudad2)"> -->
<!--     Get Distance -->
<!--   </button> -->
<!-- </template> -->
<!---->
<!-- <script> -->
<!-- export default { -->
<!--   data() { -->
<!--     return { -->
<!--       ciudad1: "", -->
<!--       ciudad2: "", -->
<!--     }; -->
<!--   }, -->
<!--   methods: { -->
<!--     setPlace1(place) { -->
<!--       console.log("setPlace1"); -->
<!--       this.ciudad1 = place.name; -->
<!--       console.log(this.ciudad1); -->
<!--     }, -->
<!--     setPlace2(place) { -->
<!--       console.log("setPlace2"); -->
<!--       this.ciudad2 = place.name; -->
<!--       console.log(this.ciudad2); -->
<!--     }, -->
<!--   }, -->
<!-- }; -->
<!-- </script> -->

<template>
  <GMapAutocomplete
    placeholder="This is a placeholder"
    @place_changed="setPlace"
  >
  </GMapAutocomplete>
</template>
<script>
export default {
  name: "App",
  data() {
    return {};
  },
  methods: {
    setPlace() {
      console.log("setPlace");
    },
  },
};
</script>
