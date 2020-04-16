package com.example.todolist.ui

import androidx.compose.Composable
import androidx.compose.state
import androidx.compose.unaryPlus
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.*
import androidx.ui.material.MaterialTheme
import androidx.ui.tooling.preview.Preview
import com.example.todolist.model.ActivitiesModel
import com.example.todolist.data.activities

@Preview
@Composable
fun App() {
    MaterialTheme(
        colors = lightThemeColors,
        typography = themeTypography
    ) {

        val activitiesModel = +state { ActivitiesModel(activities) }
        VerticalScroller {
            Column {
                activitiesModel.value.items.forEach {
                    Activity(
                        text = it.name,
                        onCopy = { activitiesModel.value.add(it.name) },
                        onRemove = { activitiesModel.value.remove(it.id) }
                    )
                }
            }
        }
    }
}
