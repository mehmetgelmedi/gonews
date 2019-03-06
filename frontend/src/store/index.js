import Vue from 'vue'
import Vuex from 'vuex'
import Request from 'request'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        thxNewsApi: true,
        isSpinloading: false,
        allNews: [],
        news: [],
        limit: 0,
        search: "technology"
    },
    mutations: {
        spinLoader(state, payload) {
            state.isSpinloading = payload.status;
        },
        thankYouNewsApi(state, payload) {
            state.thxNewsApi = payload.status;
        }
    },
    actions: {
        getNews({ state, commit, dispatch }) {
            commit('spinLoader', { status: true });
            Request(
                `http://localhost:3456/news/${state.search}`,
                (error, res, body) => {
                    const result = JSON.parse(body);
                    if (result.status === "error") {
                        commit('spinLoader', { status: false });
                        alert("Not found");
                    } else {
                        state.allNews = result;
                        commit('thankYouNewsApi', { status: false });
                        dispatch('incLimit');
                    }
                }
            );
        },
        incLimit({ state, commit }) {
            if (state.limit < state.allNews.length) {
                commit('spinLoader', { status: true });
                setTimeout(() => {
                    state.news = [
                        ...state.news,
                        ...state.allNews.slice(state.limit, (state.limit += 10))
                    ];
                    commit('spinLoader', { status: false });
                }, 2000);
            }
        },
    }
});