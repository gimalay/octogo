package com.example.todolist.ui.home

import androidx.compose.*
import androidx.ui.foundation.Clickable
import androidx.ui.layout.*
import androidx.ui.material.*
import com.example.todolist.model.ViewModelOuterClass.ViewModel
import com.example.todolist.services.Navigator
import com.example.todolist.ui.lightThemeColors
import com.example.todolist.ui.themeTypography
import com.google.protobuf.ByteString

@Composable
fun Home(
    projects: List<ViewModel.Home.Project>,
    onAddProject: (name: String) -> Unit = {},
    onRemoveProject: (id: ByteString) -> Unit = {},
    onNavigateTo: (screen: Navigator.Screen) -> Unit = {}
) {

    Column {
        HomeHeader(
            onAddProject = onAddProject
        )
        HomeBody {
            projects.forEach {
                Clickable(
                    onClick = {
                        onNavigateTo(
                            Navigator.Screen.Article(it)
                        )
                    }
                ) {
                    Article(
                        text = it.name,
                        onCopy = { onAddProject(it.name) },
                        onRemove = { onRemoveProject(it.id) }
                    )
                }
            }
        }
    }
}
