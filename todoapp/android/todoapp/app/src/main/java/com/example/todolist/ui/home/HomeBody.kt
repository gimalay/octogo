package com.example.todolist.ui.home

import androidx.compose.Composable
import androidx.compose.unaryPlus
import androidx.ui.core.Text
import androidx.ui.core.dp
import androidx.ui.foundation.VerticalScroller
import androidx.ui.layout.*
import androidx.ui.material.Button
import androidx.ui.material.ContainedButtonStyle
import androidx.ui.material.MaterialTheme
import androidx.ui.material.TextButtonStyle
import androidx.ui.tooling.preview.Preview
import com.example.todolist.model.ViewModelOuterClass
import com.example.todolist.utils.newUUIDasByteString

@Composable
fun HomeBody(
    children: @Composable() () -> Unit
) {
    VerticalScroller {
        Column {
            children()
        }
    }
}


@Composable
fun Article(
    text: String,
    onCopy: () -> Unit,
    onRemove: () -> Unit
) {
    Row(
        modifier = Spacing(left = 16.dp, right = 16.dp, top = 5.dp, bottom = 5.dp)
    ) {
        Text(
            text = text,
            modifier = Flexible(1f) wraps Gravity.Center /*wraps Spacing(16.dp)*/,
            style = (+MaterialTheme.typography()).subtitle1
        )

        Row {
            Button(
                "Copy",
                style = ContainedButtonStyle(),
                onClick = onCopy
            )
            WidthSpacer(width = 16.dp)
            Button(
                "Remove",
                style = TextButtonStyle(),
                onClick = onRemove
            )
        }
    }
}

@Preview
@Composable
fun PreviewHomeBody() {
    val projects = listOf(
        ViewModelOuterClass.ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 1")
            .build(),
        ViewModelOuterClass.ViewModel.Home.Project
            .newBuilder()
            .setID(newUUIDasByteString())
            .setName("Test project 2")
            .build()
    )
    HomeBody {
        projects.forEach {
            Article(
                text = it.name,
                onCopy = {},
                onRemove = {}
            )
        }
    }
}
