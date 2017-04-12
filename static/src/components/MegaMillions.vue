<template lang="html">
  <el-row>
    <el-collapse accordion>
      <el-collapse-item title="Search Winners" name="1">
        <search-box ballLabel="Mega Ball" multiplier="Megaplier" ballProp="mega_ball" :api="api"></search-box>
        <br>
      </el-collapse-item>
      <el-collapse-item title="Generate Picks" name="2">
        <pick-box ballLabel="Mega Ball"></pick-box>
      </el-collapse-item>
      <el-collapse-item title="Lottery Info" name="3">
            <info-box
              :howLink="howLink"
              :whereLink="whereLink"
              :faq="faq">
            </info-box>
      </el-collapse-item>
    </el-collapse>
  </el-row>
</template>

<script>
import SearchBox from './SearchBox.vue'
import PickBox from './PickBox.vue'
import InfoBox from './InfoBox.vue'

export default {
  components: {
    SearchBox,
    PickBox,
    InfoBox
  },
  data() {
    return {
      api: 'http://45.79.150.247:3000/api/megamillions',
      howLink: 'http://www.megamillions.com/how-to-play',
      whereLink: 'http://www.megamillions.com/where-to-play',
      faq: 'http://www.megamillions.com/faqs'
    }
  }
  // computer: {
  //   api: function() {
  //     if (process.env.NODE_ENV === 'production') {
  //       return 'http://localhost:3000/api/megamillions'
  //     }
  //   }
  // }
}
</script>

<style lang="css">

</style>
