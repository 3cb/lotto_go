<template lang="html">
  <div>
    <search-box ballLabel="Mega Ball" multiplier="Megaplier" ballProp="mega_ball" :api="api"></search-box>
  </div>
</template>

<script>
import SearchBox from './SearchBox.vue'

export default {
  components: {
    SearchBox
  },
  data() {
    return {
      api: 'http://localhost:8080/api/megamillions'
    }
  }
}
</script>

<style lang="css">

</style>
