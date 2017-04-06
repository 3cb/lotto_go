<template>
    <el-row>
    <el-collapse accordion>
      <el-collapse-item title="Search Winners" name="1">
        <search-box ballLabel="Powerball" multiplier="PowerPlay" ballProp="powerball" :api="api"></search-box>
        <br>
    </el-collapse-item>
    <el-collapse-item title="Generate Picks" name="2">
        <pick-box ballLabel="Powerball"></pick-box>
    </el-collapse-item>
    </el-collapse>
  </el-row>
</template>

<script>
import SearchBox from './SearchBox.vue'
import PickBox from './PickBox.vue'

export default {
    components: {
        SearchBox,
        PickBox
    },
    data() {
        return {
            api: 'http://localhost:8080/api/powerball'
        }
    }
}
</script>

<style lang="css">

</style>