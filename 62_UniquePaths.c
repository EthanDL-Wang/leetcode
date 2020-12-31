/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：5.9 MB, 在所有 Go 提交中击败了5.09%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/



int uniquePaths(int m, int n){
    int **buff = NULL;
    int *tmp = NULL;
    int i,j;
    int count = 0;
    
    buff = (int **)malloc(sizeof(int *)*m);
    if(buff == NULL){
        printf("malloc error!\n");
        return -1;
    }
    for(i = 0; i < m; i++){
        tmp = NULL;
        tmp = (int *)malloc(sizeof(int)*n);
        if(tmp == NULL){
            printf("malloc error!\n");
            return -1;
        }
        buff[i] = tmp;
    }
    
    for(i = 0; i < m; i++){
        buff[i][0] = 1;
    }
    for(i = 0; i < n; i++){
        buff[0][i] = 1;
    }    
    for(i = 1; i < m; i++){
        for(j = 1; j < n; j++){
            buff[i][j] = buff[i-1][j] + buff[i][j-1];
        }       
    }
    
    count = buff[m-1][n-1];
    for(i = 0; i < m; i++){
        free(buff[i]);
    }
    free(buff);
    return count;
}















