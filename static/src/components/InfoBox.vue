<template>
    <section>
        <h1>Some useful links:</h1>
        <ul style="list-style-type: none">
            <li><el-button @click="gotoHow">How To Play</el-button></li>
            <li><el-button @click="gotoWhere">Where To Play</el-button></li>
            <li><el-button @click="gotoFAQ">FAQ</el-button></li>
        </ul>
    </section>
</template>

<script>
export default {
    props: ['howLink', 'whereLink', 'faq'],
    methods: {
        gotoHow() {
            window.open(this.howLink)
        },
        gotoWhere() {
            window.open(this.whereLink)
        },
        gotoFAQ() {
            window.open(this.faq)
        }
    }
}
</script>

<style>
</style>