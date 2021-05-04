const env = {
    namespaced: true,
    state:{
        env:0
    },
    getters:{
        getEnv(state) {
            state.env = process.env.NODE_ENV
            return state.env
        }
    },
    mutations:{

    },
    actions:{

    }
}

export default env
