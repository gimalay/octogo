package com.example.todolist.ui

import androidx.compose.*
import androidx.ui.animation.Crossfade
import androidx.ui.material.MaterialTheme
import com.example.todolist.services.Navigator

@Composable
fun App(
    screen: Navigator.Screen,
    children: @Composable() (Navigator.Screen) -> Unit
) {
    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        Crossfade(screen) { children(it) }
    }
}
