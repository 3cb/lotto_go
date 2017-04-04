<template>
    <div>
        <search-box ball="Powerball" multiplier="PowerPlay" :api="api"></search-box>
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
            api: 'http://localhost:8080/api/powerball'
        }
    }
}
</script>

<style lang="css">

</style>