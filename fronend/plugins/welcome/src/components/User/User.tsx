import React, { FC, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme,  } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
  Container,
  Grid,
  TextField,
  Button,
  FormControl,
  MenuItem,
  Select,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntGender } from '../../api';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
    textAlign: 'right',
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface user {
  idemployee: string;
  name: string;
  gender: number;

}
 

  const User: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [gerders,setGerder] = React.useState<EntGender[]>([]);
  const [user, setUser] = React.useState<Partial<user>>({});
  const [IdemployeeError,setIdemployeeError] = React.useState('');


  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getGender = async () => {   //ดึงข้อมูล
    const res = await http.listGender({ limit: 2, offset: 0 });
    setGerder(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGender();
   }, []);

  // set data to object user
  const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
    const name = event.target.name as keyof typeof User;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setUser({ ...user, [name]: value });
    console.log(user);
  };


   const ValidateIdemployee = (val: string) => {
     return val.length == 6 ? true : false;
   }


  const checkPattern = (id: string, value: string) => {
     switch (id) {
       case 'idemployee':
           ValidateIdemployee(value) ? setIdemployeeError ('') : setIdemployeeError ("กรอกรหัสพนักงาน 6 ตัว");
         return;
       default:
         return;
     }
   }

   const alertMessage = (icon: any, title: any) => {
     Toast.fire({
       icon: icon,
      title: title,
    });
   }

   const checkCaseSaveError = (field: string) => {
     switch (field) {
       case 'idemployee':
         alertMessage("error", "กรอกรหัสพนักงาน 6 ตัว");
         return;
       default:
         alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ")
         return;


    }
 }

    // clear input form
    function clear() {
      setUser({});
    }

  
  const save = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/users';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
    };
    console.log(user);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          // Toast.fire(
          //   {
          //     icon: 'error',
          //     title: 'บันทึกไม่สำเร็จ',
          //   }
          // )
           checkCaseSaveError(data.error.Name)
        }
      });
  };

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`User`}>
        
      </Header>
      <Content>
        <Container maxWidth="md">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error={IdemployeeError ? true : false}
                className={classes.formControl}
                name = "idemployee"
                label="รหัสพนักงาน"
                variant="outlined"
                type = "string"
                value={user.idemployee || ''}
                helperText={IdemployeeError}
                onChange={handleChange}
              />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ - สกุล</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                className={classes.formControl}
                name ="name"
                type ="string"
                label="ชื่อ - สกุล"
                value={user.name || ''}
                variant="outlined"
                onChange={handleChange}
              />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
              <Select
                name="gender"
                label="เพศ"
                value={user.gender || ''}
                onChange={handleChange}
              >
                {gerders.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.gender} </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึก
              </Button>

              <Button style={{ marginLeft: 10 }}
                onClick={clear}
                variant="contained"
                color="secondary">
                CLEAR
             </Button>

              <Button style={{ marginLeft: 10 }}
                component={RouterLink} to="/Tables"
                variant="contained">
                แสดงผล
             </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default User;

