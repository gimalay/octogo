package com.example.todolist.ui

import androidx.compose.*
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.Column
import androidx.ui.material.MaterialTheme
import androidx.ui.tooling.preview.Preview
import com.example.todolist.model.HomeModelBase
import com.example.todolist.model.HomeModelMock

@Preview
@Composable
fun AppPreview() {
    val homeModel = +state { HomeModelMock() }
    App(
        home = homeModel.value
    )
}

@Composable
fun App(home: HomeModelBase) {
    +onActive {
        home.load()
    }

    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {
        VerticalScroller {
            Column {
                home.data.projectsList.forEach {
                    Activity(
                        text = it.name,
                        onCopy = { home.addProject(it.name) },
                        onRemove = { home.removeProject(it.id) }
                    )
                }
            }
        }
    }
}
