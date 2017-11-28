<template>
  <div id="app">
    <header class="header"><h1>Lotto</h1></header>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">Mega Millions</span>
      </div>
      <div>
        <mega-millions></mega-millions>
      </div>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">Powerball</span>
      </div>
      <div>
        <powerball></powerball>
      </div>
    </el-card>
  </div>
</template>

<script>
import MegaMillions from './components/MegaMillions.vue'
import Powerball from './components/Powerball.vue'

export default {
  name: 'app',
  components: {
    MegaMillions,
    Powerball
  },
  data () {
    return {
      activeName: 'first'
    }
  }
}
</script>

<style>
#app {
  position: absolute;
  top: 10%;
  left: 25%;
  width: 50%;
}
.header {
  width: 100%;
  padding: 0%;
  margin: 0%;
}

.box-card {
  margin: 0px 0px 10px 0px;
}

</style>
