package com.example.todolist.ui.common

import androidx.compose.Composable
import androidx.compose.unaryPlus
import androidx.ui.core.*
import androidx.ui.foundation.Dialog
import androidx.ui.layout.*
import androidx.ui.material.*
import androidx.ui.tooling.preview.Preview

@Composable
fun Loadable(
    show: Boolean = true,
    children: @Composable() () -> Unit
) {
    if (show) {
        Dialog(onCloseRequest = { }) {
            Container {
                Column {
                    val textStyle = (+MaterialTheme.typography()).h6
                    CurrentTextStyleProvider(textStyle) {
                        Text("Loading ...")
                    }
                }
            }
        }
    }
    else {
        children()
    }
}

@Preview
@Composable
fun PreviewLoading() {
    Loadable(show = true) {
        Text("test content")
    }
}
