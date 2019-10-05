<template>
    <nav class="navbar">
        <div class="navbar-brand">
            <nuxt-link class="navbar-item subtitle is-4" to="/">
                Youtube Player
            </nuxt-link>
        </div>

        <div class="navbar-menu">
            <div class="navbar-start">
                <div class="navbar-item">
                    <div class="field has-addons">
                        <form action="/search" method="get">
                            <input name="q" class="input" type="text" placeholder="検索">
                        </form>
                        <p class="control">
                            <button class="button">
                                Search
                            </button>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</template>
