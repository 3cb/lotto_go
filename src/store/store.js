import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        latestMM: 'megaNumber',
        latestPB: 'powerNumber'
    },
    getters: {
        getLatestMM: state => {
            return state.latestMM
        },
        getLatestPB: state => {
            return state.latestPB
        }
    },
    actions: {},
    mutations: {}
})